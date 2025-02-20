import * as esbuild from "esbuild"

const buildMode = "--build"
const watchMode = "--watch"

const helpString = `Mode must be provided as one of ${buildMode} or ${watchMode}`

const buildOptions: esbuild.BuildOptions = {
  entryPoints: {
    index: "src/index.tsx",
  },
  bundle: true,
  external: ["node:crypto"],
  write: true,
  publicPath: "/",
  sourcemap: false,
  minify: true,
  outdir: "build",
  banner: {
    js: `
      // FastestTextEncoderPolyfill
      if(typeof TextEncoder==='undefined'){var FastestTextEncoderPolyfill=function(){};FastestTextEncoderPolyfill.prototype.encode=function(str){var Len=str.length,resPos=0;var resArr=new Uint8Array(str.length*3);for(var point=0,nextcode=0,i=0;i!==Len;){point=str.charCodeAt(i),i+=1;if(point>=0xD800&&point<=0xDBFF){if(i===Len){resArr[resPos++]=0xef;resArr[resPos++]=0xbf;resArr[resPos++]=0xbd;break}nextcode=str.charCodeAt(i);if(nextcode>=0xDC00&&nextcode<=0xDFFF){point=(point-0xD800)*0x400+nextcode-0xDC00+0x10000;i+=1;if(point>0xffff){resArr[resPos++]=point>>18|240;resArr[resPos++]=point>>12&63|128;resArr[resPos++]=point>>6&63|128;resArr[resPos++]=point&63|128;continue}}}if(point<=0x007f){resArr[resPos++]=point&0xff}else if(point<=0x07ff){resArr[resPos++]=point>>6|0xc0;resArr[resPos++]=point&0x3f|0x80}else{resArr[resPos++]=point>>12|0xe0;resArr[resPos++]=point>>6&0x3f|0x80;resArr[resPos++]=point&0x3f|0x80}}return new Uint8Array(resArr.buffer.slice(0,resPos))};window.TextEncoder=FastestTextEncoderPolyfill;};
    `,
  },
}

const args = process.argv.splice(2)

if (args.length !== 1) {
  console.error(helpString)
  process.exit(1)
}

const mode = args[0]

switch (mode) {
  case buildMode:
    esbuild.build(buildOptions).catch(() => process.exit(1))
    break
  case watchMode:
    esbuild
      .context({ ...buildOptions })
      .then((context) => context.watch())
      .catch(() => process.exit(1))
    break
  default:
    console.error(helpString)
}
