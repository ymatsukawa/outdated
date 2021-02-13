class NotesController < ApplicationController
  include Authentication
  before_action :set_note, only: [:show, :edit, :update, :destroy]
  before_action :set_tag,  only: [:show, :edit]

  # GET /notes
  def index
    @notes = []
    if params[:body].present?
      @notes = Note.where('body LIKE ?', "%#{params[:body]}%").order(:created_at)
    end

    name = params[:name]
    if name.present?
      @notes = current_user.notes.select { |note| note.tags.find { |tag| tag.name == name } }
    end

    if params[:body].blank? and params[:name].blank?
      @notes = User.wrote_notes(current_user.id)
    end
    
    @tags = Tag.all.map { |tag| [tag.name, tag.name] }
  end

  # GET /notes/1
  def show
    unless UserNote.written_by_user?(params[:id], current_user.id)
      redirect_to notes_path
    end

    extensions = { fenced_code_blocks: true }
    markdown = Redcarpet::Markdown.new(Redcarpet::Render::HTML.new, extensions)
    @body = markdown.render(@note.body).html_safe
  end

  # GET /notes/new
  def new
    @note = Note.new
  end

  # GET /notes/1/edit
  def edit
    @new_tag = Tag.new
  end

  # POST /notes
  def create
    @note = Note.new(note_params)

    if @note.save and UserNote.new({ user_id: current_user.id, note_id: @note.id }).save
      redirect_to @note, notice: 'Note was successfully created.'
    else
      render :new
    end
  end

  # PATCH/PUT /notes/1
  def update
    if @note.update(note_params)
      redirect_to @note, notice: 'Note was successfully updated.'
    else
      render :edit
    end
  end

  # DELETE /notes/1
  def destroy
    @note.destroy
    redirect_to notes_url, notice: 'Note was successfully destroyed.'
  end

  # GET /notes/search
  def search

    render
  end

  private

  def set_note
    @note = Note.find(params[:id])
  end

  def set_tag
    @tags = @note.tags
  end

  # Never trust parameters from the scary internet, only allow the white list through.
  def note_params
    params.require(:note).permit(:title, :body)
  end
end
