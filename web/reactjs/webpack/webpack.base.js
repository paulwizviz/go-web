// Copyright 2019 Paul Sitoh
// SPDX-License-Identifier: Apache-2.0

const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require('copy-webpack-plugin');

const BUILD_DIR = path.resolve(__dirname, '..', 'public');
const SRC_DIR = path.resolve(__dirname, '..', 'src');

module.exports = {
  entry: path.resolve(SRC_DIR, 'index.js'),
  output: {
    path: BUILD_DIR,
    publicPath: '/',
    filename: 'bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: 'babel-loader',
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({ 
      template: path.resolve(SRC_DIR, 'index.html')
    }),
    new CopyPlugin([
      { from: path.resolve(SRC_DIR,'images'), to: path.resolve(BUILD_DIR,'images') },
    ]),
  ]
}
