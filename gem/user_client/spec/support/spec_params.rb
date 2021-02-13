USER_ID                  = 1
EMAIL                    = 'dummy@dummy.com'
PASSWORD                 = 'dummy_password'
NAME                     = 'dummy_name'
NEW_NAME                 = 'dummy_new_name'
NEW_EMAIL                = 'dummy_new@dummy.com'
NEW_PASSWORD             = 'dummy_new_password'

INVALID_USER_ID          = '123invalid_id'
INVALID_EMAIL            = 'invalid@@dummy.com'
INVALID_NAME             = '名前'
INVALID_PASSWORD         = 'パスワード'
NOT_USED_EMAIL           = 'not_used@dummy.com'
HUGE_USER_ID             = 100000
NOT_PERMITTED_OBJECTS    = [[], {}, 3.14, Object, :symbol]
NOT_PERMITTED_IDENTIFIER = [INVALID_USER_ID, INVALID_EMAIL]
NOT_PERMITTED_PROPERTY   = [:id, :updated_at, :created_at, :password_confirmation, :hack]
NOT_PERMITTED_ENTITY     = { password: INVALID_PASSWORD,
  name:     INVALID_NAME,
  email:    INVALID_EMAIL,
}

GET_ID_REQUEST    = USER_ID
GET_EMAIL_REQUEST = EMAIL

PUT_IDENTIIFIER = USER_ID
PUT_PROPERTY    = :name
PUT_ENTITY      = NAME

POST_REQUEST = {
  email:    NEW_EMAIL,
  name:     NEW_NAME,
  password: NEW_PASSWORD,
}

DELETE_REQUEST_ID    = USER_ID
DELETE_REQUEST_EMAIL = EMAIL

MODEL_USER_ATTRIBUTES = {
  email:    EMAIL,
  name:     NAME,
  password: PASSWORD,
}
