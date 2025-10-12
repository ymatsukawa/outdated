require "sinatra"

not_found do
  erb :not_found
end

error 400 do
  erb :bad_request
end

get "/example/html" do
  @first_name = params[:first_name]
  @last_name = params[:last_name]
  erb :html_index
end

post "/example/machine" do
  nut = params[:nut]
  iron_plate = params[:iron_plate]

  if nut.nil? || iron_plate.nil?
    return 400
  else
    status 201
    erb :html_machine_created
  end
end

put "/example/clothes" do
  necktie = params[:necktie]
  suit = params[:suit]

  if necktie.nil? || suit.nil?
    return 400
  else
    status 200
    erb :html_clothes_updated
  end
end

delete "/example/paper" do
  paper_id = params[:paper_id]

  if paper_id.nil?
    return 400
  else
    status 200
    erb :html_paper_deleted
  end
end