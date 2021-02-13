
shared_examples_for 'have a hashnize DB entity' do
  let(:body) { subject[:body] }

  [:body, 'body'].each do |root_property|
    it { expect(subject).to have_key(root_property) }
  end

  PERMITTED_PROPERTY.each do |property|
    it { expect(body).to have_key(property) }
    it { expect(body[property]).not_to be_an_instance_of(Hash) }
  end
end
