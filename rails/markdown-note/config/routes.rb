Rails.application.routes.draw do
  get 'welcome/index', to: 'welcome#index'
  devise_for :users, controllers: { omniauth_callbacks: 'users/omniauth_callbacks' }
  devise_scope :user do
    get 'login', to: 'devise/sessions#new'
  end

  resources :notes do
    namespace :notes do
      resources :tags, only: [:create, :destroy]
    end
  end

  get :search, to: 'notes#index'
  get :search_tag, to: 'notes#index'
end
