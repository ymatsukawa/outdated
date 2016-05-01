# added to enable stub login at request spec @ custom_v0.0.1

shared_context "hook 'before and after' process" do
  before { Warden.test_mode!  }
  after  { Warden.test_reset! }
end

shared_context 'prepare stub sign_in/sign_out methods' do
  def sign_in(user)
    login_as(user, scope: :user)
  end

  def sign_out
    logout(:user)
  end
end
