const path = require('path');

module.exports = {
  mode: 'production',
  entry: './src/go-app/.gae/static/js/main.js',
  output: {
    path: path.resolve(__dirname, 'src/go-app/.gae/static/'),
    filename: '[name].min.js'
  },
};
