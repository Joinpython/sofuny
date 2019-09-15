
import Vue from 'vue'
import hljs from 'highlight.js'
import 'highlight.js/styles/atom-one-dark.css'

Vue.directive("highlight", function (el) {
  let blocks = el.querySelectorAll("pre code");
  blocks.forEach((item)=>{
    hljs.highlightBlock(item)
  })
});
