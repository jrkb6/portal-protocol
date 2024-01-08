import{s as u,n as i,o as y}from"../chunks/scheduler.d9671216.js";import{S as f,i as b,e as o,a as c,f as d,g as x,h as n,x as h,k as w}from"../chunks/index.30f8fe88.js";import{p as v}from"../chunks/stores.b8b10ded.js";function m(l){let t,p='<div class="grid grid-cols-3 gap-4 mb-6"><div><h3 class="text-green-400">BNB</h3> <p>$317.70</p></div> <div><h3 class="text-blue-400">New Listing</h3> <p>NFP $1.04</p></div> <div><h3 class="text-yellow-400">Top Gainer Coin</h3> <p>FARM $53.82</p></div></div> <div class="mb-6"><ul class="flex space-x-2"><li class="bg-blue-500 px-4 py-2 rounded">Favorites</li> <li class="bg-blue-500 px-4 py-2 rounded">All Cryptos</li> <li class="bg-blue-500 px-4 py-2 rounded">My Wallet</li></ul></div> <div class="overflow-x-auto"><table class="min-w-full divide-y divide-gray-700"><thead><tr><th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider">Name</th> <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider">Price</th> <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider">24h</th> <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider">Change</th> <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider">24h Volume</th> <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider">Market Cap</th> <th class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider">Actions</th></tr></thead> <tbody class="divide-y divide-gray-700"><tr><td class="px-6 py-4 whitespace-nowrap">BTC Bitcoin</td> <td class="px-6 py-4 whitespace-nowrap">$42,254.89</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-0.61%</td> <td class="px-6 py-4 whitespace-nowrap">$26.08B</td> <td class="px-6 py-4 whitespace-nowrap">$827.46B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">ETH Ethereum</td> <td class="px-6 py-4 whitespace-nowrap">$2,302.83</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-1.62%</td> <td class="px-6 py-4 whitespace-nowrap">$16.43B</td> <td class="px-6 py-4 whitespace-nowrap">$276.76B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">AVAX Avalanche</td> <td class="px-6 py-4 whitespace-nowrap">$2,304.09</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-0.61%</td> <td class="px-6 py-4 whitespace-nowrap">$26.08B</td> <td class="px-6 py-4 whitespace-nowrap">$827.46B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">XRP Ripple</td> <td class="px-6 py-4 whitespace-nowrap">$2,102.13</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-1.62%</td> <td class="px-6 py-4 whitespace-nowrap">$10.43B</td> <td class="px-6 py-4 whitespace-nowrap">$123.76B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">DOT Polkadot</td> <td class="px-6 py-4 whitespace-nowrap">$42,254.89</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-0.61%</td> <td class="px-6 py-4 whitespace-nowrap">$22.08B</td> <td class="px-6 py-4 whitespace-nowrap">$153.46B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">DOGE DogeCoin</td> <td class="px-6 py-4 whitespace-nowrap">$2,302.83</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-1.62%</td> <td class="px-6 py-4 whitespace-nowrap">$23.43B</td> <td class="px-6 py-4 whitespace-nowrap">$11.12B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">ADA Cardano</td> <td class="px-6 py-4 whitespace-nowrap">$2,254.89</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-0.61%</td> <td class="px-6 py-4 whitespace-nowrap">$6.08B</td> <td class="px-6 py-4 whitespace-nowrap">$27.46B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">ETH Ethereum</td> <td class="px-6 py-4 whitespace-nowrap">$2,302.83</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-1.62%</td> <td class="px-6 py-4 whitespace-nowrap">$16.43B</td> <td class="px-6 py-4 whitespace-nowrap">$276.76B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">BTC Bitcoin</td> <td class="px-6 py-4 whitespace-nowrap">$42,254.89</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-0.61%</td> <td class="px-6 py-4 whitespace-nowrap">$26.08B</td> <td class="px-6 py-4 whitespace-nowrap">$827.46B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr> <tr><td class="px-6 py-4 whitespace-nowrap">USDT USD Tether</td> <td class="px-6 py-4 whitespace-nowrap">$1,102.11</td> <td class="px-6 py-4 whitespace-nowrap text-red-500">-1.62%</td> <td class="px-6 py-4 whitespace-nowrap">$10.23B</td> <td class="px-6 py-4 whitespace-nowrap">$7.76B</td> <td class="px-6 py-4 whitespace-nowrap"><a href="#" class="text-blue-500 hover:text-blue-600">Detail</a> <a href="#" class="text-blue-500 hover:text-blue-600 ml-4">Trade</a></td></tr></tbody></table></div>';return{c(){t=x("div"),t.innerHTML=p,this.h()},l(e){t=n(e,"DIV",{class:!0,"data-svelte-h":!0}),h(t)!=="svelte-1ags8vr"&&(t.innerHTML=p),this.h()},h(){w(t,"class","bg-gray-900 text-white p-6")},m(e,a){c(e,t,a)},d(e){e&&d(t)}}}function g(l){let t,p='<h1 class="text-2xl font-bold">Your fancy De-Fi. Please connect your wallet.</h1> <header class="container mx-auto my-20 text-center"><h2 class="text-5xl font-bold mb-6">Shaping web that will captivate the world.</h2> <p class="mb-6">To access your account, please verify that you are over 18 years of age.</p></header> <div class="flex flex-wrap justify-center gap-4"><div class="bg-purple-700 shadow-lg rounded-lg p-8 flex flex-col items-center justify-center"><p class="text-3xl font-semibold">30+</p> <p>Currencies</p></div> <div class="bg-purple-700 shadow-lg rounded-lg p-8 flex flex-col items-center justify-center"><p class="text-3xl font-semibold">30,000+</p> <p>Trades</p></div> <div class="bg-purple-700 shadow-lg rounded-lg p-8 flex flex-col items-center justify-center"><p class="text-3xl font-semibold">$150,000+</p> <p>Total Assets</p></div></div>';return{c(){t=x("div"),t.innerHTML=p,this.h()},l(e){t=n(e,"DIV",{class:!0,"data-svelte-h":!0}),h(t)!=="svelte-11wj4bv"&&(t.innerHTML=p),this.h()},h(){w(t,"class","bg-gray-900 text-white p-6 grid h-screen place-items-center")},m(e,a){c(e,t,a)},d(e){e&&d(t)}}}function $(l){let t;function p(s,r){if(!s[0])return g;if(s[0])return m}let e=p(l),a=e&&e(l);return{c(){a&&a.c(),t=o()},l(s){a&&a.l(s),t=o()},m(s,r){a&&a.m(s,r),c(s,t,r)},p(s,[r]){e!==(e=p(s))&&(a&&a.d(1),a=e&&e(s),a&&(a.c(),a.m(t.parentNode,t)))},i,o:i,d(s){s&&d(t),a&&a.d(s)}}}function B(l,t,p){let e;return y(()=>{const a=v.subscribe(s=>{p(0,e=s)});return()=>{a()}}),[e]}class k extends f{constructor(t){super(),b(this,t,B,$,u,{})}}export{k as component};
