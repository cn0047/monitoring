function d(elId, title, rows) {
  const data = new google.visualization.DataTable();
  data.addColumn('date', '');
  data.addColumn('number', '');
  data.addRows(rows);
  const options = {
    legend: {position: 'none'},
    title: title,
  };
  const chart = new google.charts.Line(document.getElementById(elId));
  chart.draw(data, google.charts.Line.convertOptions(options));
}

const f = function (prj, limit) {
  fetch(`/api/charts?project=${prj}&limit=${limit}`)
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
  f('realtimelog-health-check', 20)
};

google.charts.load('current', {packages: ['line']});
google.charts.setOnLoadCallback(app);
