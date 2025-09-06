import { createServer } from "node:http";
import { toNodeListener } from "h3";
import { app } from "./app";
import process from "node:process";

const port = process.env.port || 3000;

createServer(toNodeListener(app))
.listen(port, () => {
    console.info(`🎯 Server started on port ${port} !`);
});
