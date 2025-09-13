import { useFetch } from '@vueuse/core'

export class HttpService {
  BASE_URL: string = '/api'

  constructor() {}

  async get<T>({
    path,
    params = {},
  }: {
    path: string
    params?: Record<string, string>
  }): Promise<T> {
    const queryParams = params.length ? '?' + new URLSearchParams(params).toString() : ''

    const { data, error } = await useFetch(this.BASE_URL + path + queryParams, {
      method: 'GET',
    }).json()

    if (error.value) {
      console.error(`Error during a GET request ${path} :`, error.value)
    }

    return data.value as T
  }

  async post<T>({ path, body }: { path: string; body: object }): Promise<T> {
    const { data, error } = await useFetch(this.BASE_URL + path, {
      method: 'POST',
      body: JSON.stringify(body),
    }).json()

    if (error.value) {
      console.error(`Error during a POST request ${path} :`, error.value)
    }

    return data.value as T
  }

  async patch<T>({ path, body }: { path: string; body: object }): Promise<T> {
    const { data, error } = await useFetch(this.BASE_URL + path, {
      method: 'PATCH',
      body: JSON.stringify(body)
    }).json();

    if (error.value) {
      console.error(`Error during a PATCH request ${path} :`, error.value)
    }

    return data.value as T;
  }
}
