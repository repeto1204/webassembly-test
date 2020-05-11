const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
    mode: "development",
    entry: './src/index.tsx',
    output: {
        path: __dirname + '/build/',
        filename: '[name].js'
    },
    module: {
        rules:  [
            {
                test: /\.(ts|tsx)$/,
                exclude: '/node_modules/',
                use: {
                    loader: "babel-loader"
                },
            }
        ]
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js']
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: "./public/index.html"
        })
    ],
    devServer: {
        host: "localhost",
        port: "1234",
        contentBase: './public',
        watchContentBase: true,
        hot: true,
        overlay: true
    }
}