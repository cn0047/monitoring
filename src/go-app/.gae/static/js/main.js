/**
 * Shows element by id.
 *
 * @param {String} id Element id.
 */
const show = function (id) {
  document.getElementById(id).style.display = 'block';
};

/**
 * Hides element by id.
 *
 * @param {String} id Element id.
 */
const hide = function (id) {
  document.getElementById(id).style.display = 'none';
};

/**
 * Returns true in case when provided parameter is object.
 *
 * @param {Object} o Object to inspect.
 *
 * @returns {boolean} True in case when provided parameter is object.
 */
const isObject = function (o) {
  return typeof o === 'object' && Array.isArray(o) === false && o !== null;
};

/**
 * Performs chart related error rendering.
 *
 * @param {String} error Error message.
 */
const renderChartError = function (error) {
  document.querySelector('#chartError p').innerHTML = error;
  show('chartError');
};

/**
 * Performs ChartRT rendering.
 *
 * @param {Array} rows Array with chart rows.
 * @param {String} title Chart title.
 */
const renderChartRT = function (rows, title) {
  if (rows.length === 0) {
    document.getElementById('chartRT').innerHTML = 'No data.';
    return;
  }

  const dt = new google.visualization.DataTable();
  dt.addColumn('date', title);
  dt.addColumn('number', '');
  dt.addRows(rows);
  const options = {
    legend: {position: 'none'},
    title: '',
  };
  const chart = new google.charts.Line(document.getElementById('chartRT'));
  chart.draw(dt, google.charts.Line.convertOptions(options));
};

/**
 * Performs ChartRC rendering.
 *
 * @param {Array} rows Array with chart rows.
 * @param {String} title Chart title.
 */
const renderChartRC = function (rows, title) {
  if (rows.length === 0) {
    document.getElementById('chartRC').innerHTML = 'No data.';
    return;
  }

  const dt = new google.visualization.DataTable();
  dt.addColumn('date', title);
  dt.addColumn('number', '');
  dt.addRows(rows);
  const options = {
    legend: {position: 'none'},
    title: '',
  };
  const chart = new google.charts.Line(document.getElementById('chartRC'));
  chart.draw(dt, google.charts.Line.convertOptions(options));
};

/**
 * Performs charts's payload processing.
 *
 * @param {Object} payload Chart's payload data from server.
 */
const handleChartsPayload = function (payload) {
  if (payload.error !== null) {
    renderChartError(payload.error.data);
    hide('loader');
    return;
  }
  hide('chartError');

  const data = payload.success.data;
  data.rt.map((el) => {el[0] = new Date(el[0])});
  data.rc.map((el) => {el[0] = new Date(el[0])});

  renderChartRT(data.rt, data.titleRT);
  renderChartRC(data.rc, data.titleRC);
  hide('loader');
};

/**
 * Performs charts rendering.
 *
 * Facade function which performs:
 * fetch charts data related to current project from server
 * and render charts data or error message.
 *
 * @param {String} timeRange Time range string.
 * @param {Number} limit Limit value.
 */
const renderCharts = function (timeRange, limit) {
  show('loader');
  const prj = document.getElementById('selectedProject').value;
  fetch(`/api/v1/charts?project=${prj}&timeRange=${timeRange}&limit=${limit}`)
    .then(res => res.json())
    .then(payload => handleChartsPayload(payload));
};

/**
 * Performs project's errors rendering.
 *
 * @param {Object|string} errors Errors object or error message.
 */
const renderProjectFormErrors = function (errors) {
  const ul = document.querySelector('#projectFormErrors ul');
  ul.innerHTML = '';

  if (isObject(errors)) {
    for (key in errors) {
      const li = document.createElement('li');
      li.innerHTML = errors[key];
      ul.appendChild(li);
    }
  } else {
    const li = document.createElement('li');
    li.innerHTML = errors;
    ul.appendChild(li);
  }

  hide('projectFormSuccess');
  show('projectFormErrors');
};

/**
 * Adds new project name into dropdown with all available projects.
 *
 * @param {string} name project name.
 */
const addProjectIntoSelect = function (name) {
  const o = document.createElement('option');
  o.value = name;
  o.text = name;

  const s = document.getElementById('selectedProject');
  s.appendChild(o);
  s.value = name;
};

/**
 * Performs payload processing for project form.
 *
 * @param {String} name Project's name.
 * @param {Object} payload Project's payload data from server.
 */
const handleProjectForm = function (name, payload) {
  if (payload.error !== null) {
    renderProjectFormErrors(payload.error.data);
    return;
  }

  hide('projectFormErrors');
  show('projectFormSuccess');
  addProjectIntoSelect(name);
  setTimeout(() => {
    hide('projectFormSuccess');
    hide('projectForm');
    show('charts');
  }, 3000);
};

/**
 * Performs project form save.
 *
 * Facade function which performs:
 * send request to server to create new project
 * and perform response payload processing.
 */
const submitProjectForm = function () {
  const form = {
    name: document.getElementById('name').value,
    url: document.getElementById('url').value,
    schedule: document.getElementById('schedule').value,
    method: document.getElementById('method').value,
    json: document.getElementById('json').value,
  };
  const args = {
    method: 'post', headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(form)
  };
  fetch('/api/v1/projects', args)
    .then(res => res.json())
    .then(payload => handleProjectForm(form.name, payload));
};

/**
 * Gets current "time range" value.
 *
 * @returns {String} Current "time range" value.
 */
const getCurrentTimeRange = function () {
  return document.querySelector('#charts .timeRangeActive').text;
};

/**
 * Application entry point.
 */
const app = function () {
  if (initError !== "") {
    renderChartError(initError);
    return;
  }

  renderCharts(getCurrentTimeRange(), 0);
};

/**
 * Handler for dropdown with projects, which performs charts re-rendering.
 */
document.getElementById('selectedProject').addEventListener('change', () => {
  renderCharts(getCurrentTimeRange(), 0);
});

/**
 * Handler for "time range" quick links.
 */
document.querySelectorAll('#charts .timeRange a').forEach((el) => {
  el.addEventListener('click', (e) => {
    document.querySelectorAll('#charts .timeRange a').forEach((aEl) => {
      aEl.className = 'timeRangeDefault';
    });
    e.target.className = 'timeRangeActive';
    renderCharts(e.target.text, 0);
  });
});

/**
 * Handler to open project's form.
 */
document.getElementById('addProject').addEventListener('click', () => {
  hide('charts');
  show('projectForm');
});

/**
 * Handler to close project's form.
 */
document.getElementById('projectFormBtnCancel').addEventListener('click', () => {
  hide('projectForm');
  hide('projectFormErrors');
  hide('projectFormSuccess');
  show('charts');
});

/**
 * Handler to save project's form.
 */
document.getElementById('projectFormBtnSave').addEventListener('click', submitProjectForm);



google.charts.load('current', {packages: ['line']});
google.charts.setOnLoadCallback(app);
