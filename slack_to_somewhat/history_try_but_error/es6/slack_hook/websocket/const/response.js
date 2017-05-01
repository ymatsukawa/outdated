export const STATUS_CODE = {
  SUCCESS: 200,
  ERROR: 500
}

export const successRes = (message = 'success') => {
  return {
    statusCode: STATUS_CODE.SUCCESS,
    message: message
  }
}

export const errorRes = (message = 'error') => {
  return {
    statusCode: STATUS_CODE.ERROR,
    message: message
  }
}

