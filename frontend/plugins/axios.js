import pkg from '../../config'

export default function ({ $axios, redirect, store}) {
  let auth = {
    username: pkg.basicAuth.username,
    password: pkg.basicAuth.password,
    key: pkg.basicAuth.key
  };
  $axios.setHeader("Authorization", 'Bearer '+pkg.basicAuth.key);
  $axios.onRequest(config => {
    console.log('Making <' + config.method +'> request to ' + config.url)
  });

  $axios.onResponse(response => {
    // console.log(response);
    // store.dispatch("acToken", response.headers["set-cookie"][0])
  });

  $axios.onError(error => {
    // console.log(error.response.status)
    // console.log(error.response.statusText)
    // console.log(error.response.headers)
    // console.log(error.response.config)
    const code = parseInt(error.response && error.response.status);
    if (code === 400) {
      redirect('/400')
    }
  })
}
