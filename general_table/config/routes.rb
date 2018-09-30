Rails.application.routes.draw do
  resources :tables, only: [:index, :new, :create]

  namespace :management do
    resources :base_tables, param: :base_table_id do
      member do
        resources :custom_fields, param: :custom_field_id
      end
    end
  end
end
