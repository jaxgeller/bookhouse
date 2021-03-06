import path from 'path'
import webpack from 'webpack'

export default {
  context: path.join(__dirname, './src'),

  resolve: {
    extensions: ['', '.js', '.jsx', '.scss']
  },

  entry: './index.jsx',

  output: {
    filename: 'bundle.js',
    path: path.join(__dirname, './build')
  },

  module: {
    loaders: [
      { test: /\.jsx?$/, loaders: ['babel'], exclude: /node_modules/ }
    ]
  }
}
