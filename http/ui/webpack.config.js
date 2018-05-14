const path = require('path');

module.exports = {
    context: path.join(__dirname, 'src'),
    entry: [
        './main.js',
    ],
    output: {
        path: path.join(__dirname, 'html'),
        filename: 'bundle.js',
    },
    mode: "none",
    watch: true,
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: [
                    'babel-loader'
                ],
            },
            {
                test:/\.css$/,
                use:['style-loader','css-loader']
            }
        ],
    },
};