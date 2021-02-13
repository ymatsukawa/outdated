Rails.application.routes.draw do
  namespace :backlog do
    resource :ticket, only: :create
    resource :info, only: :show
  end

  get  '*not_found' => 'application#routing_error'
  post '*not_found' => 'application#routing_error'
end
