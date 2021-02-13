require 'spec_helper'
require 'pry'

describe 'UserCilent' do
  describe 'the response of #put' do
    let(:identifier) { PUT_IDENTIIFIER }
    let(:property)   { PUT_PROPERTY }
    let(:entity)     { PUT_ENTITY }

    subject { User.new.put(identifier, { property => entity }) }

    context 'with valid request' do
      before { create_user }
      after  { UserClient::Model::User.delete_all }

      [USER_ID, EMAIL].each do |identifier|
        { password: PASSWORD, email: EMAIL }.each do |property, entity|
          context "with identifier #{identifier} and #{property} as #{entity}" do
            let(:identifier) { identifier }
            let(:property)   { property }
            let(:entity)     { entity }

            it { expect(subject[:body][property]).to eq entity }
            it_behaves_like 'have a hashnize DB entity'
          end
        end
      end
    end

    context 'with invalid request' do
      [HUGE_USER_ID, NOT_USED_EMAIL].each do |identifier|
        context "with identifier as #{identifier}" do
          let(:identifier) { identifier }

          it { expect { subject }.to raise_error UserClient::BadRequest::ContentError }
        end
      end

      NOT_PERMITTED_IDENTIFIER.each do |identifier|
        context "with identifier as #{identifier}" do
          let(:identifier) { identifier }

          it { expect { subject }.to raise_error UserClient::BadRequest::IdentifierError }
        end
      end

      NOT_PERMITTED_PROPERTY.each do |property|
        context "with property as #{property}" do
          let(:property) { property }

          it { expect { subject }.to raise_error UserClient::BadRequest::PropertyError }
        end
      end

      [:name, :password].each do |property|
        context 'with invalid name or password' do
          before { create_user }
          after  { delete_user }
          let(:property) { property }
          let(:entity)   { 'ぜんかく' }

          it { expect { subject }.to raise_error UserClient::BadRequest::ContentError }
        end
      end
    end
  end
end
