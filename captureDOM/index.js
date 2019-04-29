const settings = require('./settings/settings');
const Authn    = require('./domain/authn');
const Capture  = require('./domain/capture');
const Transfer = require('./domain/transfer');

const authn = new Authn(settings.authn, settings.request);
try {
  authn.authenticate((cookies) => {
    const capture = new Capture(settings.captures, settings.io, settings.request, cookies);
    capture.capture(() => {
      const transfer = new Transfer(settings.io);
      transfer.transferCaptures();
    });
  });
} catch(e) {
  console.log(e.message);
  return; // exit
}
