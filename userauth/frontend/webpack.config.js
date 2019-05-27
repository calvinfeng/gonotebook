module.exports = {
    mode: "development",
    devtool: "inline-source-map",
    entry: "./src/index.tsx",
    output: {
        path: __dirname + "/../public/assets/javascripts",
        filename: "index.bundle.js"
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js", ".jsx"]
    },
    module: {
        rules: [
            { test: /\.tsx?$/, loader: "ts-loader" },
            { test: /\.scss$/, use: ["style-loader", "css-loader", "sass-loader"] }
        ]
    }
};