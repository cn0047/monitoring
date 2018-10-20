const DEFAULT_TIME_RANGE = '12h';

const toggle = function (id) {
  const s = document.getElementById(id).style;
  const d = s.display;
  s.display = (d === '' || d === 'none') ? 'block' : 'none';
};

const show = function (id) {
  document.getElementById(id).style.display = 'block';
};

const hide = function (id) {
  document.getElementById(id).style.display = 'none';
};

const renderChartError = function (error) {
  document.querySelector('#chartError p').innerHTML = error;
  show('chartError');
};

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

const fillCharts = function (payload) {
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

const renderCharts = function (timeRange, limit) {
  show('loader');
  const prj = document.getElementById('selectedProject').value;
  fetch(`/api/v1/charts?project=${prj}&timeRange=${timeRange}&limit=${limit}`)
    .then(res => res.json())
    .then(payload => fillCharts(payload));
};

const isObject = function (o) {
  return typeof o === 'object' && Array.isArray(o) === false && o !== null;
};

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

const addProjectIntoSelect = function (name) {
  const o = document.createElement('option');
  o.value = name;
  o.text = name;

  const s = document.getElementById('selectedProject');
  s.appendChild(o);
  s.value = name;
};

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

const app = function () {
  renderCharts(DEFAULT_TIME_RANGE, 0);
};

google.charts.load('current', {packages: ['line']});
google.charts.setOnLoadCallback(app);

document.getElementById('selectedProject').addEventListener('click', () => {
  renderCharts(DEFAULT_TIME_RANGE, 0);
});

document.getElementById('selectedProject').addEventListener('change', () => {
  renderCharts(DEFAULT_TIME_RANGE, 0);
});

document.querySelectorAll('#charts .timeRange a').forEach((el) => {
  el.addEventListener('click', (e) => renderCharts(e.target.text, 0));
});

document.getElementById('addProject').addEventListener('click', () => {
  hide('charts');
  show('projectForm');
});

document.getElementById('projectFormBtnCancel').addEventListener('click', () => {
  hide('projectForm');
  hide('projectFormErrors');
  hide('projectFormSuccess');
  show('charts');
});

document.getElementById('projectFormBtnSave').addEventListener('click', submitProjectForm);
