require 'spec_helper'
require 'active_support'

describe 'UserCilent' do
  let(:finder) { { id: USER_ID } }

  describe 'the response of #get' do
    let(:request) { GET_ID_REQUEST }

    subject { User.new.get(request) }

    context 'with valid identifier request' do
      before { create_user }
      after  { UserClient::Model::User.delete_all }

      [GET_ID_REQUEST, GET_EMAIL_REQUEST].each do |id|
        let(:request) { id }

        it_behaves_like 'have a hashnize DB entity'
      end

      [HUGE_USER_ID, NOT_USED_EMAIL].each do |id|
        context 'with fake id' do
          let(:request) { id }

          it { expect { subject }.to raise_error UserClient::BadRequest::ContentError }
        end
      end
    end

    context 'with invalid identifier request' do
      NOT_PERMITTED_IDENTIFIER.each do |id|
        context "with identifier as #{id}" do
          let(:request) { id }

          it { expect { subject }.to raise_error UserClient::BadRequest::IdentifierError }
        end
      end

      context 'with not permitted object' do
        NOT_PERMITTED_OBJECTS.each do |id|
          let(:request) { id }

          it { expect { subject }.to raise_error UserClient::BadRequest::TypeError }
        end
      end
    end
  end
end
