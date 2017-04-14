var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
    entry: {
        'js/bundle.js': './resources/js/app.js',
        'css/main.css': './resources/css/app.scss'
    },
    output: {
        filename: './public/[name]'
    },
    devtool: process.env.NODE_ENV !== 'production' ? 'source-map' : false,
    module: {
        loaders: [
            {
                test: /\.jsx?$/,
                exclude: /(node_modules|bower_components)/,
                loader: 'babel-loader',
                query: {
                    presets: ['es2015']
                }
            },
            {
                test: /\.scss$/,
                include: /resources\/css/,
                loaders: ExtractTextPlugin.extract('css-loader!sass-loader')
            },
            {
                test: /\.(eot|svg|ttf|woff|woff2)$/,
                loader: 'file-loader?name=public/fonts/[name].[ext]'
            }
        ]
    },
    plugins: [
        new ExtractTextPlugin({
            filename: './public/css/main.css',
            allChunks: true
        })
    ]
};