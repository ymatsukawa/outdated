# added to commonize test data @ custom_v0.0.1

SCHEME = 'http'
HOST   = 'localhost'
PORT   = '3000'
ROOT_URI = SCHEME + '://' + HOST + ':' + PORT

NAME                   = 'dummy_name'
EMAIL                  = 'example@example.com'
CURRENT_PASSWORD       = 'password'
PASSWORD               = 'new_password'
PASSWORD_CONFIRMATION  = PASSWORD
NOTICE_MAIL_SEND_STATE = 0

TITLE              = 'dummy_title'
BODY               = <<CONTENT
#{Settings.test.notify.notice_email.article.body}
CONTENT
TAG_LIST           = 'dummy, tag, list'
IS_PUBLIC_EDITABLE = 1
LOCK_VERSION       = 0 # mutex > variable is const 0

COMMENT_BODY = <<CONTENT
#{Settings.test.notify.notice_email.comment.body}
CONTENT

BLANK      = ''
MINUS_NUM  = -1

INVALID_USER_ID               = MINUS_NUM
INVALID_NAME                  = BLANK
INVALID_EMAIL                 = BLANK
INVALID_CURRENT_PASSWORD      = 'invalid_current_password'
INVALID_PASSWORD              = 'invalid_password'
INVALID_PASSWORD_CONFIRMATION = 'invalid_password_confirmation'
INVALID_MAIL_SEND_STATE       = MINUS_NUM

INVALID_TITLE              = BLANK
INVALID_BODY               = BLANK
INVALID_IS_PUBLIC_EDITABLE = MINUS_NUM
INVALID_LOCK_VERSION       = MINUS_NUM

VALID_USERS_RESOURCE = {
  name:                   NAME,
  email:                  EMAIL,
  current_password:       CURRENT_PASSWORD,
  notice_mail_send_state: NOTICE_MAIL_SEND_STATE,
  password:               PASSWORD,
  password_confirmation:  PASSWORD_CONFIRMATION,
}

VALID_ARTICLE_RESOURCE = {
  title:              TITLE,
  body:               BODY,
  tag_list:           TAG_LIST,
  is_public_editable: IS_PUBLIC_EDITABLE,
  lock_version:       LOCK_VERSION,
  user_id:            'pass resource (user.id or like 1) in test case',
}

VALID_USERS_REQUEST_PARAMS = {
  user: VALID_USERS_RESOURCE
}

INVALID_USERS_REQUEST_PARAMS_LIST = {
  name:                   INVALID_NAME,
  email:                  INVALID_EMAIL,
  notice_mail_send_state: INVALID_MAIL_SEND_STATE,
  current_password:       INVALID_CURRENT_PASSWORD,
  password:               INVALID_PASSWORD,
  password_confirmation:  INVALID_PASSWORD_CONFIRMATION,
}

VALID_ARTICLE_REQUEST_PARAMS = {
  article: VALID_ARTICLE_RESOURCE
}

INVALID_ARTICLE_REQUEST_PARAMS_LIST = {
  title: INVALID_TITLE,
  body:  INVALID_BODY,
  # below is success case, check after customize
  ## is_public_editable: INVALID_IS_PUBLIC_EDITABLE,
  ## lock_version:       INVALID_LOCK_VERSION,
  ## user_id:            INVALID_USER_ID,
}

VALID_ARTICLE_CONTROLLER_PARAMS        = VALID_ARTICLE_REQUEST_PARAMS
INVALID_ARTICLE_CONTROLLER_PARAMS_LIST = INVALID_ARTICLE_REQUEST_PARAMS_LIST
ARTICLE_SHOW_PATH_FORMAT               = show_path_of('articles')

VALID_COMMENTS_RESOURCE = {
  body:       COMMENT_BODY,
  user_id:    'pass resource (user.id or like 1) in test case',
  article_id: 'pass resource (article.id or like 1) at test',
}

VALID_COMMENTS_REQUEST_PARAMS = {
  comment: VALID_COMMENTS_RESOURCE
}

VALID_COMMENTS_CONTROLLER_PARAMS = VALID_COMMENTS_REQUEST_PARAMS
