require 'rails_helper'

describe 'Article', type: :model do
  using RSpec::Parameterized::TableSyntax
  
  describe '#only_visible' do
    context 'when archived record exists' do
      let(:record_size) { 2 }
      before { create_list(:article, record_size) }

      it do 
        expect(Article.only_visible.size).to eq record_size
      end
    end

    context 'when archived record does not exist' do
      let(:non_archive_size) { 3 }
      let(:archived_size) { 2 }
      before do
        create_list(:article, non_archive_size)
        create_list(:article, archived_size, :archive)
      end

      it do 
        expect(Article.only_visible.size).to eq non_archive_size
      end
    end
  end

  describe '#create_record' do
    let(:params) { { title: title, text: text } }
    subject { Article.create_record(params) }

    context 'when params are valid' do
      let(:title) { 'this is title' }
      let(:text) { 'Lorem ipsum dolor sit amet' }
      
      it { is_expected.to be_truthy }
    end

    context 'when params are invalid' do
      where(:title_, :text_) do
        nil | 'a'
        'a' | nil
      end
      with_them do
        let(:title) { title_ }
        let(:text) { text_ }
  
        it { is_expected.to be_falsey }
      end
    end
  end

  describe '#archive' do
    subject { Article.archive(0) }

    context 'when id exist' do
      before { create(:article) }
      let(:id) { Article.first }

      subject { Article.archive(id) }
      it { is_expected.to be_truthy }
    end

    context 'when id does not exist' do
      let(:id) { 0 }
      it { is_expected.to be_falsey }
    end
  end
end