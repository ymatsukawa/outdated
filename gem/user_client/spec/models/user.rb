require 'spec_helper'
require 'active_record'

describe UserClient::Models::User do
  describe 'format' do
    before { create_user }
    after  { delete_user }
    let(:attributes) { MODEL_USER_ATTRIBUTES }

    subject { UserClient::Models::User.new(attributes).save! }

    context 'with valid attributes' do
      it { expect { subject }.not_to raise_error }

      context 'within maximum length attributes' do
        let(:attributes) do
          MODEL_USER_ATTRIBUTES.merge({
            email:    'a'.rjust(245) + '@dummy.com',
            name:     'a'.rjust(255),
            password: 'a'.rjust(255),
          })
        end

        it { expect { subject }.not_to raise_error }
      end

      context 'within minimum length attributes' do
        let(:attributes) { MODEL_USER_ATTRIBUTES.merge(password: 'a'.rjust(8)) }

        it { expect { subject }.not_to raise_error }
      end
    end

    context 'with invalid attributes' do
      ['', 'a'.rjust(246) + '@dummy.com'].each do |email|
        context 'with invalid email' do
          let(:attributes) { MODEL_USER_ATTRIBUTES.merge(email: email) }

          it { expect { subject }.to raise_error ActiveRecord::RecordInvalid }
        end
      end

      ['', 'a'.rjust(256), 'なまえ'].each do |name|
        context 'with invalid name' do
          let(:attributes) { MODEL_USER_ATTRIBUTES.merge(name: name) }

          it { expect { subject }.to raise_error ActiveRecord::RecordInvalid }
        end
      end

      ['', 'a'.rjust(7), 'a'.rjust(256), 'パスワード'].each do |password|
        context 'with invalid password' do
          let(:attributes) { MODEL_USER_ATTRIBUTES.merge(password: password) }

          it { expect { subject }.to raise_error ActiveRecord::RecordInvalid }
        end
      end
    end
  end
end
