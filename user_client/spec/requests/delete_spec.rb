require 'spec_helper'

describe '#delete' do
  let(:finder) { DELETE_REQUEST_ID }

  describe 'the response of #create' do
    let(:request) { DELETE_REQUEST_ID }

    subject { User.new.delete(request) }

    context 'with valid request' do
      before { create_user }

      { id: DELETE_REQUEST_ID, email: DELETE_REQUEST_EMAIL }.each do |property, entity|
        context 'with valid request' do
          let(:request) { entity }
          let(:finder)  { { property => entity } }

          it { expect(subject).to be true }
        end
      end
    end

    context 'with invalid request' do
      [HUGE_USER_ID, NOT_USED_EMAIL].each do |id|
        context "with identifier as #{id}" do
          let(:request) { id }

          it { expect { subject }.to raise_error UserClient::BadRequest::ContentError }
        end
      end

      NOT_PERMITTED_OBJECTS.each do |id|
        context 'with not permitted id' do
          let(:request) { id }

          it { expect { subject }.to raise_error UserClient::BadRequest::TypeError }
        end
      end
    end
  end
end
