function d(elId, title, rows) {
  const data = new google.visualization.DataTable();
  data.addColumn('date', title);
  data.addColumn('number', '');
  data.addRows(rows);
  const options = {
    legend: {position: 'none'},
    title: '',
  };
  const chart = new google.charts.Line(document.getElementById(elId));
  chart.draw(data, google.charts.Line.convertOptions(options));
}

const f = function (timeRange, limit) {
  const prj = document.getElementById('project').value;
  fetch(`/api/charts?project=${prj}&timeRange=${timeRange}&limit=${limit}`)
    .then(res => res.json())
    .then(payload => {
      const data = payload.success.data;
      data.rt.map((el) => {el[0] = new Date(el[0])});
      data.rc.map((el) => {el[0] = new Date(el[0])});
      d('chartRT', data.rtTitle, data.rt);
      d('chartRC', data.rcTitle, data.rc);
    })
};

const app = function () {
  f('12h', 0)
};

google.charts.load('current', {packages: ['line']});
google.charts.setOnLoadCallback(app);

document.getElementById('project').addEventListener('click', () => f('12h', 0));

document.querySelectorAll('#charts .timeRange a').forEach((el) => {
  el.addEventListener('click', (e) => f(e.target.text, 0));
});

document.getElementById('addProject').addEventListener('click', () => {
  document.getElementById('addForm').style.display = 'block';
  document.getElementById('charts').style.display = 'none';
});

document.getElementById('btnCancel').addEventListener('click', () => {
  document.getElementById('addForm').style.display = 'none';
  document.getElementById('charts').style.display = 'block';
});

document.getElementById('btnSave').addEventListener('click', () => {
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
  fetch('/api/projects', args)
    .then(res => res.json())
    .then(payload => {
      console.log(payload)
    });
});
