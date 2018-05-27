
module.exports = {
    mode: "development",
    devtool: "inline-source-map",
    entry: "./ui/index.tsx",
    output: {
        path: __dirname + "/public",
        filename: "index.bundle.js"
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js"]
    },
    module: {
        rules: [
            { test: /\.tsx?$/, loader: "ts-loader" }
        ]
    }
};