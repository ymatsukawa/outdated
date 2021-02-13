class Backlog::TicketsController < ApplicationController
  def create
    render json: Backlog::Ticket.new_instance(ticket_params).create
  end

  def ticket_params
    params.permit(C.api.backlog.request.permitted_keys)
  end
end
