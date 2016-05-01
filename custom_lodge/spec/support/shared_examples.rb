# added to commonize test result @ custom_v0.0.1

shared_context 'prepare shared given' do
  Given(:status_code)  { res.status }
end

shared_examples_for 'return response status code with' do |expect_status_code|
  include_context 'prepare shared given'

  Then { expect(status_code).to eq expect_status_code }
end

shared_examples_for 'validate to specific field' do |factory_name, field, values, expected_validation|
  subject(:model)      { FactoryGirl.create(factory_name) }
  subject(:validation) { model.send(:validate) }

  values.each do |value|
    context "when passed #{field}'s as #{value}" do
      before { model.send("#{field}=", value) }
      it     { expect(validation).to eq expected_validation }
    end
  end
end
