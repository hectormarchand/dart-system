export class StreamService {

    private framesBufferMap: Record<string, Buffer>

    constructor() {
        this.framesBufferMap = {
            leftFrame: Buffer.alloc(0),
            frontFrame: Buffer.alloc(0),
            rightFrame: Buffer.alloc(0)
        };
    }

    saveCameraFrames(leftFrameBuffer: Buffer, frontFrameBuffer: Buffer, rightFrameBuffer: Buffer) {
        this.framesBufferMap.leftFrame = leftFrameBuffer;
        this.framesBufferMap.frontFrame = frontFrameBuffer;
        this.framesBufferMap.rightFrame = rightFrameBuffer;
    }

    getLeftCameraFrame(): Buffer {
        return this.framesBufferMap.leftFrame;
    }

    getFrontCameraFrame(): Buffer {
        return this.framesBufferMap.frontFrame;
    }

    getRightCameraFrame(): Buffer {
        return this.framesBufferMap.rightFrame;
    }


}