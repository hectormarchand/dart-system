import {defineEventHandler, readFormData, readMultipartFormData, Router} from "h3";
import {StreamService} from "./stream.service";

export class StreamWebservice {

    private readonly ROUTE_PREFIX = "/stream";

    private streamService: StreamService;

    constructor(streamService: StreamService, apiRouter: Router) {
        this.streamService = streamService;
        this.initRoutes(apiRouter);
    }

    initRoutes(router: Router) {
        router.get(
            this.ROUTE_PREFIX + "/left-camera",
            defineEventHandler(event => {
                const res = event.node.res;
                const req = event.node.req;
                res.setHeader('Content-Type', 'multipart/x-mixed-replace; boundary=--frame');
                res.setHeader('Cache-Control', 'no-cache');
                res.setHeader('Connection', 'close');

                const writeFrame = () => {
                    const buffer = this.streamService.getLeftCameraFrame();

                    if (buffer && buffer.length > 0) {
                        res.write(`--frame\r\n`);
                        res.write(`Content-Type: image/jpeg\r\n`);
                        res.write(`Content-Length: ${buffer.length}\r\n\r\n`);
                        res.write(buffer);
                        res.write('\r\n');
                    }
                };

                writeFrame();
                const interval = setInterval(writeFrame, 50); // ~10 fps

                req.on('close', () => {
                    console.log("closing")
                    clearInterval(interval);
                });
            })
        );

        router.get(
            this.ROUTE_PREFIX + "/front-camera",
            defineEventHandler(event => {
                const res = event.node.res;
                const req = event.node.req;
                res.setHeader('Content-Type', 'multipart/x-mixed-replace; boundary=--frame');
                res.setHeader('Cache-Control', 'no-cache');
                res.setHeader('Connection', 'close');

                const writeFrame = () => {
                    const buffer = this.streamService.getFrontCameraFrame();

                    if (buffer && buffer.length > 0) {
                        res.write(`--frame\r\n`);
                        res.write(`Content-Type: image/jpeg\r\n`);
                        res.write(`Content-Length: ${buffer.length}\r\n\r\n`);
                        res.write(buffer);
                        res.write('\r\n');
                    }
                };

                writeFrame();
                const interval = setInterval(writeFrame, 50); // ~10 fps

                req.on('close', () => {
                    console.log("closing")
                    clearInterval(interval);
                });
            })
        );

        router.get(
            this.ROUTE_PREFIX + "/right-camera",
            defineEventHandler(event => {
                const res = event.node.res;
                const req = event.node.req;
                res.setHeader('Content-Type', 'multipart/x-mixed-replace; boundary=--frame');
                res.setHeader('Cache-Control', 'no-cache');
                res.setHeader('Connection', 'close');

                const writeFrame = () => {
                    const buffer = this.streamService.getRightCameraFrame();

                    if (buffer && buffer.length > 0) {
                        res.write(`--frame\r\n`);
                        res.write(`Content-Type: image/jpeg\r\n`);
                        res.write(`Content-Length: ${buffer.length}\r\n\r\n`);
                        res.write(buffer);
                        res.write('\r\n');
                    }
                };

                writeFrame();
                const interval = setInterval(writeFrame, 50); // ~10 fps

                req.on('close', () => {
                    console.log("closing")
                    clearInterval(interval);
                });
            })
        );

        router.post(
            this.ROUTE_PREFIX + "/cameras",
            defineEventHandler(async event => {
                const files = await readMultipartFormData(event) as { name: string, type: string, data: Buffer }[];

                if (files.length !== 3) {
                    throw new Error("The number of frames received is not equal to three, actual number : " + files.length);
                }

                this.streamService.saveCameraFrames(files[0].data, files[1].data, files[2].data);

                return { status: "ok" };
            })
        )
    }

    streamMjpegOverHttp(res: any, req: any, getBufferFunction: Function) {
        res.setHeader('Content-Type', 'multipart/x-mixed-replace; boundary=frame');
        res.setHeader('Cache-Control', 'no-cache');
        res.setHeader('Connection', 'close');

        const writeFrame = () => {
            const buffer = getBufferFunction();

            if (buffer && buffer.length > 0) {
                res.write(`--frame\r\n`);
                res.write(`Content-Type: image/jpeg\r\n`);
                res.write(`Content-Length: ${buffer.length}\r\n\r\n`);
                res.write(buffer);
                res.write('\r\n');
            }
        };

        writeFrame();
        const interval = setInterval(writeFrame, 100); // ~10 fps

        req.on('close', () => {
            clearInterval(interval);
        });
    }
}