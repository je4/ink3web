package static

import "embed"

// go:embed pdfjs/lib/*
//
// go:embed templates/default-book-view.html
// go:embed sounds/*
// go:embed images/dark-loader.gif images/light-loader.gif
//
//go:embed js/search.js js/d3.js js/d3bubble.js
//go:embed bootstrap/css/bootstrap.min.css bootstrap/css/bootstrap.min.css.map
//go:embed bootstrap/js/bootstrap.bundle.min.js bootstrap/js/bootstrap.bundle.min.js.map
//go:embed bootstrap-icons/font/bootstrap-icons.min.css bootstrap-icons/font/fonts/bootstrap-icons.woff2
//go:embed bootstrap-icons/font/fonts/bootstrap-icons.woff
//go:embed css/ibm-plex-mono.css css/ibm-plex-sans-condensed.css css/ibm-plex-sans.css css/ibm-plex-serif.css
//go:embed img/*
//go:embed flag-icons/css/flag-icons.min.css flag-icons/flags/4x3/de.svg flag-icons/flags/4x3/gb-eng.svg
//go:embed flag-icons/flags/4x3/fr.svg flag-icons/flags/4x3/it.svg
//go:embed videojs/video-js.min.css videojs/video.min.js
//go:embed pdf.js/web/viewer.css pdf.js/web/viewer.html pdf.js/web/viewer.mjs pdf.js/web/images/* pdf.js/web/locale/*
//go:embed pdf.js/build/pdf.mjs pdf.js/web/cmaps/* pdf.js/build/pdf.worker.mjs
//go:embed openseadragon/images/* openseadragon/openseadragon.min.js openseadragon/openseadragon.min.js.map
//go:embed css/_csp_dark.scss css/_csp_light.scss css/csp.css css/ibm-plex*.css
//go:embed css/_ink_dark.scss css/_csp_light.scss css/ink.css
//go:embed dflip/css/* dflip/js/* dflip/fonts/* dflip/images/* dflip/sound/*
//go:embed foliatereader/reader.js foliate-js/*.js foliate-js/ui/* foliate-js/vendor/*
//go:embed performance_manifest.json performance_*.png
var FS embed.FS
