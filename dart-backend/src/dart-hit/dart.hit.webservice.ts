import {DartHitService} from "./dart.hit.service";
import {defineEventHandler, readBody, Router} from "h3";
import {DartHitDto} from "../common/dart.dtos";

export class DartHitWebService {

    private readonly ROUTE_PREFIX = "/dart-hit"

    private dartHitService: DartHitService;

    constructor(dartHitService: DartHitService, apiRouter: Router) {
        this.dartHitService = dartHitService;
        this.initRoutes(apiRouter);
    }

    initRoutes(router: Router) {
        // Called from the opencv c++ module
        router.post(
            this.ROUTE_PREFIX,
            defineEventHandler(async event => {
                const dartHitDto: DartHitDto = await readBody<DartHitDto>(event);
                await this.dartHitService.addDartHit(dartHitDto);
                return { status: "ok" };
            })
        )

      router.get(
        this.ROUTE_PREFIX + "/stream-hits",
        defineEventHandler((event) => {
          return this.dartHitService.streamDartHits(event);
        })
      )
    }

}