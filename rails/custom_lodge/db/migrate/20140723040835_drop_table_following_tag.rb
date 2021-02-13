class DropTableFollowingTag < ActiveRecord::Migration
  def up
    drop_table :following_tags
  end
  def down
    create_table :following_tags do |t|
      t.references :user, null: false, index: true
      t.references :tag, null: false, index: true
      t.foreign_key :users, dependent: :delete

      t.timestamps
    end

    add_index :following_tags, [:user_id, :tag_id], unique: true
  end
end
