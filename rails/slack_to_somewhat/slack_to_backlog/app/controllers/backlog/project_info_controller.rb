class Backlog::ProjectInfoController < ApplicationController
  def show
  end

  def project_info_params
    params.permit(C.api.backlog.request.permitted_keys)
  end
end
