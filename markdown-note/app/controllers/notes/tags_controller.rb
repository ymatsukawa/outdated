module Notes
  class TagsController < ApplicationController
    def create
      note_id = tag_params[:note_id].to_i

      tag = Tag.find_by(name: tag_params[:name])
      if tag.blank?
        tag = Tag.new(name: tag_params[:name])
        tag.save and NoteTag.new({ note_id: note_id, tag_id: tag.id }).save
        flash[:error] = tag.errors.full_messages.to_sentence if tag.errors.present?
      else
        tag_id = NoteTag.find_by(note_id: note_id).tag_id
        NoteTag.new({ note_id: note_id, tag_id: tag.id }).save if tag_id.blank?
      end
      redirect_to edit_note_path(note_id)
    end

    def destroy
      note_id, tag_id = tag_params[:note_id], tag_params[:id]
      note_tag = NoteTag.find_by(note_id: note_id, tag_id: tag_id)
      if note_tag.present? and note_tag.destroy
        redirect_to edit_note_path(note_id)
      else
        raise
      end
    end

    private

    def tag_params
      params.permit(:id, :note_id, :name)
    end
  end
end
