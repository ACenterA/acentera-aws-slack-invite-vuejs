import axios from 'axios'
import localforage from 'localforage'
import { setupCache } from 'axios-cache-adapter'

const localCacheStore = localforage.createInstance({
  // List of drivers used
  driver: [
    localforage.INDEXEDDB,
    localforage.LOCALSTORAGE
  ],
  // Prefix all storage keys to prevent conflicts
  name: 'admin-rest-cache'
})

const cache = setupCache({
  maxAge: 15 * 60 * 1000,
  debug: true,
  localCacheStore
})

// create an axios instance
const service = axios.create({
  baseURL: process.env.BASE_API, // api 的 base_url
  timeout: 55000 // request timeout
})

const cachedService = axios.create({
  adapter: cache.adapter,
  baseURL: process.env.BASE_API, // api 的 base_url
  timeout: 55000 // request timeout
})

const configHandler = function(config) {
  // Do something before request is sent
  return config // console.error(er)
}

const handleSuccess = function(response) {
  const res = response
  if (res.status >= 200 || res.status <= 204) { // 200, 201, 204 /// all good
    return response
  } else {
    // TODO: On 409, retry ?
    if (res.status === 409) {
      // Ask the user to reload its browser, new verseion of this website exists
    }

    if (res.status === 401) {
      // return store.dispatch('LogOut').then(() => {
      //   location.reload()// In order to re-instantiate the vue-router object to avoid bugs
      // })
    }
    return Promise.reject(response)
  }
  // response => {
  //   const res = response.data
  //   if (res.code !== 20000) {
  //     Message({
  //       message: res.message,
  //       type: 'error',
  //       duration: 5 * 1000
  //     })
  //     // 50008:非法的token; 50012:其他客户端登录了;  50014:Token 过期了;
  //     if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
  //       // 请自行在引入 MessageBox
  //       // import { Message, MessageBox } from 'element-ui'
  //       MessageBox.confirm('你已被登出，可以取消继续留在该页面，或者重新登录', '确定登出', {
  //         confirmButtonText: '重新登录',
  //         cancelButtonText: '取消',
  //         type: 'warning'
  //       }).then(() => {
  //         store.dispatch('FedLogOut').then(() => {
  //           location.reload() // 为了重新实例化vue-router对象 避免bufg
  //         })
  //       })
  //     }
  //     return Promise.reject('error')
  //   } else {
  //     return response.data
  //   }
}
const handleError = function(error, test) {
  return new Promise(function(resolve, reject) {
    // const status = error.response ? error.response.status : null
    // For LogOut someone might already have remoed this session ...
    // Also sessions does auto-expires if left over on the server side or on password / role changes
    var config = null
    var resp = null
    var err = {}

    if (error) {
      if (error.config) {
        config = error.config
      } else {
        if (error.response && error.response.config) {
          config = error.response.config
        }
      }

      if (error.response) {
        resp = error.response
      } else {
        if (error.response && error.response) {
          resp = error.response
        }
      }
      err['config'] = config
      err['response'] = resp
    }
    if (!err.config) {
      err = error
    }

    // if (config && config.url.endsWith('/login/logout')) {
    if (!resp || resp.status >= 400) {
      return reject(resp)
    }
    return resolve()
  })
}

cachedService.interceptors.request.use(
  config => configHandler(config)
)
cachedService.interceptors.response.use(
  response => handleSuccess(response),
  error => handleError(error)
)

// response interceptor
service.interceptors.request.use(
  config => configHandler(config)
)

service.interceptors.response.use(
  response => handleSuccess(response),
  error => handleError(error)
)

const handler = function(data, cache) {
  if (cache) {
    data['cache'] = cache
    return cachedService(data)
  } else {
    return service(data)
  }
}
window.$request = handler
export default handler
