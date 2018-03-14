package badge

import "strings"

var FlatTemplate = strings.TrimSpace(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{dx}}" height="20">
  <linearGradient id="smooth" x2="0" y2="100%">
    <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
    <stop offset="1" stop-opacity=".1"/>
  </linearGradient>

  <mask id="round">
    <rect width="{{dx}}" height="20" rx="3" fill="#fff"/>
  </mask>

  <g mask="url(#round)">
    <rect width="{{subjectDx}}" height="20" fill="#555"/>
    <rect x="{{subjectDx}}" width="{{statusDx}}" height="20" fill="{{color}}"/>
    <rect width="{{dx}}" height="20" fill="url(#smooth)"/>
  </g>

  <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
    <text x="{{subjectX}}" y="15" fill="#010101" fill-opacity=".3">{{subject}}</text>
    <text x="{{subjectX}}" y="14">{{subject}}</text>
    <text x="{{statusX}}" y="15" fill="#010101" fill-opacity=".3">{{status}}</text>
    <text x="{{statusX}}" y="14">{{status}}</text>
  </g>
</svg>
`)
