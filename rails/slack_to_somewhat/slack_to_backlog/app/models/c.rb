class C < Settingslogic
  source "#{Rails.root}/config/constant.yml"
  namespace Rails.env
  suppress_errors Rails.env.production?
end
