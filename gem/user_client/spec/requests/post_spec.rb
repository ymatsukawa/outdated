require 'spec_helper'

describe 'UserCilent' do
  describe 'the response of #create' do
    let(:request) { POST_REQUEST }

    subject { User.new.post(request) }

    context 'with valid request' do
      after  { UserClient::Model::User.delete_all }

      it_behaves_like 'have a hashnize DB entity'
    end

    context 'with invalid request' do
      NOT_PERMITTED_PROPERTY.each do |property|
        context 'with not permitted property' do
          let(:request) { POST_REQUEST.merge(property => 'dummy_entity') }

          it { expect { subject }.to raise_error UserClient::BadRequest::PropertyError }
        end
      end

      [:email, :name, :password].each do |property|
        NOT_PERMITTED_OBJECTS.each do |obj|
          context 'with not permitted object' do
            let(:request) { POST_REQUEST.merge(property => obj) }

            it { expect { subject }.to raise_error UserClient::BadRequest::TypeError }
          end
        end
      end

      ['', 'a'.rjust(256), 'なまえ'].each do |name|
        context 'with invalid name' do
          let(:request) { POST_REQUEST.merge(name: name) }

          it { expect { subject }.to raise_error UserClient::BadRequest::ContentError }
        end
      end

      ['', 'a'.rjust(7), 'a'.rjust(256), 'パスワード'].each do |password|
        context 'with invalid password' do
          let(:request) { POST_REQUEST.merge(password: password) }

          it { expect { subject }.to raise_error UserClient::BadRequest::ContentError }
        end
      end
    end
  end
end
