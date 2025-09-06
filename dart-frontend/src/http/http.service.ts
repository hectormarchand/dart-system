export class HttpService {

    BASE_URL: string = "/api";

    constructor() {}

    async get<T>({ path, params = {} }: { path: string, params?: Record<string, string> }): Promise<T> {
        const queryParams = params.length ? "?" + new URLSearchParams(params).toString() : "";

        let json = undefined;
        try {
          const response = await fetch(this.BASE_URL + path + queryParams);
          if ("application/json" === response.headers.get("Content-Type")) {
            json = await response.json();
          }
        } catch (e) {
          console.error(`Error during a GET request ${path} :`, e);
        }

        return json as T;
    }

    async post<T>({ path, body }: { path: string, body: object }): Promise<T> {
      let json;
      try {
        const response = await fetch(this.BASE_URL + path, { method: "POST", body: JSON.stringify(body) });
        json = await response.json();
      } catch (e) {
        console.error(`Error during a POST request ${path} :`, e);
      }

      return json as T;
    }
}
