const fs = require('fs');
const woff2base64 = require('woff2base64');
const fonts = {
    'MaterialIcons-Regular.woff': fs.readFileSync('node_modules/material-icons/iconfont/MaterialIcons-Regular.woff'),
};
const options = {
    fontFamily: 'Material Icons'
};
const css = woff2base64(fonts, options);

fs.writeFileSync('src/material.icons.css', css.woff);