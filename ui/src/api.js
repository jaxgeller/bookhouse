const BASE_URL = 'http://localhost:8080/api'

const API = function(url) {
  return fetch(`${BASE_URL}${url}`)
}

export default API;
