package essaymatcher

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRandomProxyClient mocks the randomProxyClient interface.
type MockRandomProxyClient struct {
	mock.Mock
}

func (m *MockRandomProxyClient) NewHTTPClientWithRandomProxy() (*http.Client, error) {
	args := m.Called()
	return args.Get(0).(*http.Client), args.Error(1)
}

// TestScraper_Scrap tests the Scrap function of the Scraper.
func TestScraper_Scrap(t *testing.T) {
	// Create a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a mock HTML response
		w.Write([]byte(`    </script><div class=caas-body-wrapper><button class="caas-button noborder caas-body-collapse-button" data-ylk="itc:1;slk:Read full article">Read full article<i></i></button><div class=caas-body-content><div class=caas-body-inner-wrapper><div class=caas-body-section><div class=caas-content data-wf-sticky-boundary=caas-body-section data-wf-sticky-offset="165px   10px" data-wf-sticky-target=.caas-share-buttons><div class=caas-content-wrapper><header class=caas-header><div class=caas-logo></div><div class=caas-title-wrapper><h1 data-test-locator=headline>Headline test 1</h1><div class=caas-subheadline></div></div></header><div class=caas-content-byline-wrapper><div class="caas-attr author"><div class="caas-content-author-photo  "><a class="link caas-attr-logo" href=https://www.engadget.com/about/editors/jon-fingas/  data-ylk="elm:author;slk:Jon Fingas;itc:0"><img class=caas-img alt="Jon Fingas" src=https://s.yimg.com/ny/api/res/1.2/wHBTQKS1uRQ9XSMO7toiLg--/YXBwaWQ9aGlnaGxhbmRlcjt3PTgwO2g9ODA-/https://o.aolcdn.com/images/dims?image_uri=https%3A%2F%2Fo.aolcdn.com%2Fimages%2Fdims%3Fimage_uri%3Dhttps%253A%252F%252Fo.aolcdn.com%252Fimages%252Fdims%253Fimage_uri%253Dhttp%25253A%25252F%25252Fwww.blogcdn.com%25252Fwww.engadget.com%25252Fmedia%25252F2013%25252F01%25252Fjon-fingas-january-2013.jpg%2526compress%253D1%2526progressive%253D1%2526quality%253D75%2526client%253Dhawkeye%2526signature%253D634cded79dea613d81385e5a0cd907f7a375d004%26compress%3D1%26progressive%3D1%26quality%3D75%26client%3Dhawkeye%26signature%3Dc6679941fc083778f31a04cf6ae64c94e73190dc&compress=1&progressive=1&quality=75&client=hawkeye&signature=d9c6f530af8c10bb15b24f29a8df5ee2a881991d data-src=https://s.yimg.com/ny/api/res/1.2/wHBTQKS1uRQ9XSMO7toiLg--/YXBwaWQ9aGlnaGxhbmRlcjt3PTgwO2g9ODA-/https://o.aolcdn.com/images/dims?image_uri=https%3A%2F%2Fo.aolcdn.com%2Fimages%2Fdims%3Fimage_uri%3Dhttps%253A%252F%252Fo.aolcdn.com%252Fimages%252Fdims%253Fimage_uri%253Dhttp%25253A%25252F%25252Fwww.blogcdn.com%25252Fwww.engadget.com%25252Fmedia%25252F2013%25252F01%25252Fjon-fingas-january-2013.jpg%2526compress%253D1%2526progressive%253D1%2526quality%253D75%2526client%253Dhawkeye%2526signature%253D634cded79dea613d81385e5a0cd907f7a375d004%26compress%3D1%26progressive%3D1%26quality%3D75%26client%3Dhawkeye%26signature%3Dc6679941fc083778f31a04cf6ae64c94e73190dc&compress=1&progressive=1&quality=75&client=hawkeye&signature=d9c6f530af8c10bb15b24f29a8df5ee2a881991d></a></div><div class=caas-attr-meta><div class=caas-attr-item-author><span class=caas-author-byline-collapse><a class=link href=https://www.engadget.com/about/editors/jon-fingas/  data-ylk="elm:author;slk:Jon Fingas;itc:0">Jon Fingas</a></span></div><div class=caas-attr-job-title><span>Reporter</span></div><div class=caas-attr-time-style><time class=caas-attr-meta-time dateTime=2019-08-25T01:34:00.000Z>Sun, Aug 25, 2019, 4:34 AM</time><span class=caas-attr-meta-separator>&#183;</span><span class=caas-attr-mins-read>2 min read</span></div></div></div><div class=caas-content-contain-share><div class="caas-share-section "><div class=caas-share-buttons><div id=caas-consolidated-share-btn class=caas-consolidated-share-btn><button class="wafer-toggle caas-button noborder" data-wf-toggle-target=#caas-share-popup data-wf-toggle-class=click:toggle:share-button-hide data-wf-toggle-boundary=caas-content-byline-wrapper data-ylk=slk:share-menu-open;elm:share;sec:consolidated-share-popup;subsec:grade;itc:1 aria-label=Share><span class="arrowUpAndOut icon"><svg width=16 height=16 viewBox="0 0 24 24"><path d="M21 14.989a1 1 0 0 1 1 1v5.999a1 1 0 0 1-1 1.001H3.001A1 1 0 0 1 2 21.988v-5.999a1 1 0 0 1 2 0v5h16v-5a1 1 0 0 1 1-1ZM12.001 1l5.714 5.347a.959.959 0 1 1-1.357 1.357l-3.357-3.357v12.642H11V4.347L7.643 7.704a.96.96 0 0 1-1.358-1.357L12.001 1Z" /></svg></span></button><div id=caas-share-popup class="caas-share-popup share-button-hide" data-wf-toggle-target=#caas-consolidated-share-btn data-wf-toggle-class=clickOutside:add:share-button-hide data-wf-toggle-boundary=caas-content-byline-wrapper><div class="caas-share-popup-wrapper wafer-toggle" data-wf-toggle-target=#caas-share-popup data-wf-toggle-class=clickOutside:add:share-button-hide data-wf-toggle-boundary=caas-content-byline-wrapper><button class="link caas-button noborder copylink-url wafer-clipboard-copy " title="Copy link" data-ylk=t4:ctrl;sec:consolidated-share-popup;elm:share;elmt:btn;slk:copy-link;itc:0;outcm:share;rspns:op data-wf-copy-text=https://www.engadget.com/2019-08-24-crime-allegation-in-space.html? data-wf-state-copy-text=[state.copylink]><span class="copylink icon"><svg width=20 height=20 viewBox="0 0 24 24"><path d="M7.371 8.785a5.001 5.001 0 0 1 7.071 0 1 1 0 1 1-1.414 1.415 3 3 0 0 0-4.243 0l-4.83 4.83c-1.148 1.148-1.318 3.038-.23 4.245a3 3 0 0 0 4.353.117l1.414-1.414a.999.999 0 1 1 1.414 1.414l-1.414 1.414a5 5 0 0 1-7.4-.364c-1.665-2.012-1.367-5.012.48-6.859Zm6.212-6.212C15.43.727 18.43.428 20.443 2.091a5 5 0 0 1 .363 7.401l-4.95 4.95a5 5 0 0 1-7.07 0 1 1 0 0 1 1.413-1.414 3 3 0 0 0 4.243 0l4.95-4.95a3 3 0 0 0-.117-4.353c-1.207-1.088-3.097-.918-4.246.23L13.735 5.25a1 1 0 0 1-1.414 0l-.083-.094a.999.999 0 0 1 .083-1.32Z" /></svg></span></button><a class="link caas-button noborder mail" href=mailto:?subject=Divorce%20dispute%20leads%20to%20accusation%20of%20crime%20in%20space&body=https%3A%2F%2Fwww.engadget.com%2F2019-08-24-crime-allegation-in-space.html%3Fsoc_src%3Dsocial-sh%26soc_trk%3Dma rel="noopener noreferrer" target=_blank title=Email data-ylk=t4:ctrl;sec:consolidated-share-popup;elm:share;elmt:sh-ma;slk:Email;itc:0;outcm:share;rspns:op><span class="mail icon"><svg width=20 height=20 viewBox="0 0 512 512"><path d="M460.586 91.31H51.504c-10.738 0-19.46 8.72-19.46 19.477v40.088l224 104.03 224-104.03v-40.088c0-10.757-8.702-19.478-19.458-19.478M32.046 193.426V402.96c0 10.758 8.72 19.48 19.458 19.48h409.082c10.756 0 19.46-8.722 19.46-19.48V193.428l-224 102.327-224-102.327z" /></svg></span></a><a class="link caas-button noborder facebook" href=https://www.facebook.com/dialog/feed?app_id=458584288257241&link=https%3A%2F%2Fwww.engadget.com%2F2019-08-24-crime-allegation-in-space.html%3Fsoc_src%3Dsocial-sh%26soc_trk%3Dfb%26tsrc%3Dfb rel="noopener noreferrer" target=_blank title="Share on Facebook" data-ylk=t4:ctrl;sec:consolidated-share-popup;elm:share;elmt:sh-fb;slk:Facebook;itc:0;outcm:share;rspns:op><span class="facebook icon"><svg width=20 height=20 viewBox="0 0 32 32"><path d="M12.752 30.4V16.888H9.365V12.02h3.387V7.865c0-3.264 2.002-6.264 6.613-6.264 1.866 0 3.248.19 3.248.19l-.11 4.54s-1.404-.013-2.943-.013c-1.66 0-1.93.81-1.93 2.152v3.553h5.008l-.22 4.867h-4.786V30.4h-4.88z" /></svg></span></a><a class="link caas-button noborder twitter" href=https://twitter.com/intent/tweet?text=Divorce%20dispute%20leads%20to%20accusation%20of%20crime%20in%20space&url=https%3A%2F%2Fwww.engadget.com%2F2019-08-24-crime-allegation-in-space.html%3Fsoc_src%3Dsocial-sh%26soc_trk%3Dtw%26tsrc%3Dtwtr&via=engadget rel="noopener noreferrer" target=_blank title=Tweet data-ylk=t4:ctrl;sec:consolidated-share-popup;elm:share;elmt:sh-tw;slk:Twitter;itc:0;outcm:share;rspns:op><span class="twitter icon"><svg width=20 height=20 viewBox="0 0 24 24"><path d="M13.9027 10.7686L21.3482 2.3H19.5838L13.119 9.6532L7.95547 2.3H2L9.8082 13.4193L2 22.3H3.76443L10.5915 14.5348L16.0445 22.3H22L13.9023 10.7686H13.9027ZM11.4861 13.5173L10.695 12.4101L4.40018 3.59968H7.11025L12.1902 10.7099L12.9813 11.8172L19.5847 21.0594H16.8746L11.4861 13.5177V13.5173Z" /></svg></span></a><a class="link caas-button noborder whatsapp" href=https://api.whatsapp.com/send?text=https%3A%2F%2Fwww.engadget.com%2F2019-08-24-crime-allegation-in-space.html%3Fsoc_src%3Dsocial-sh%26soc_trk%3Dwa rel="noopener noreferrer" target=_blank title="Share on WhatsApp" data-ylk=t4:ctrl;sec:consolidated-share-popup;elm:share;elmt:sh-wh;slk:Whatsapp;itc:0;outcm:share;rspns:op><span class="whatsapp icon"><svg width=20 height=20 viewBox="0 0 20 20"><path d="M10.0025195,0 L9.99751953,0 C4.48376953,0 1.953125e-05,4.485 1.953125e-05,10 C1.953125e-05,12.1875 0.704980469,14.215 1.90373047,15.86125 L0.657519531,19.57625 L4.50123047,18.3475 C6.08248047,19.395 7.96876953,20 10.0025195,20 C15.5162305,20 20.0000195,15.51375 20.0000195,10 C20.0000195,4.48625 15.5162305,0 10.0025195,0 M15.8212656,14.1212461 C15.5799766,14.8024961 14.6225156,15.3674961 13.8587656,15.5324961 C13.3362656,15.6437852 12.6537266,15.7324961 10.3562656,14.7799961 C7.41751563,13.5624961 5.52501562,10.5762461 5.37751563,10.3824961 C5.23622656,10.1887461 4.19001563,8.80124609 4.19001563,7.36624609 C4.19001563,5.93124609 4.91876563,5.23249609 5.21251562,4.93249609 C5.45376562,4.68624609 5.85251562,4.57374609 6.23501562,4.57374609 C6.35876563,4.57374609 6.46997656,4.57999609 6.56997656,4.58499609 C6.86376563,4.59749609 7.01122656,4.61499609 7.20501563,5.07874609 C7.44626563,5.65999609 8.03376563,7.09499609 8.10376562,7.24249609 C8.17497656,7.38999609 8.24626562,7.58999609 8.14626562,7.78374609 C8.05251562,7.98374609 7.97001563,8.07249609 7.82247656,8.24249609 C7.67501563,8.41249609 7.53501562,8.54249609 7.38751563,8.72499609 C7.25251563,8.88374609 7.09997656,9.05374609 7.26997656,9.34749609 C7.43997656,9.63499609 8.02747656,10.5937461 8.89251563,11.3637461 C10.0087266,12.3574961 10.9137266,12.6750352 11.2374766,12.8099961 C11.4787656,12.9099961 11.7662266,12.8862461 11.9424766,12.6987461 C12.1662656,12.4574961 12.4425156,12.0574961 12.7237266,11.6637461 C12.9237266,11.3812461 13.1762656,11.3462461 13.4412266,11.4462852 C13.7112266,11.5399961 15.1400156,12.2462461 15.4337266,12.3924961 C15.7275156,12.5400352 15.9212656,12.6099961 15.9925156,12.7337461 C16.0625156,12.8574961 16.0625156,13.4387461 15.8212656,14.1212461" /></svg></span></a><a class="link caas-button noborder reddit" href="https://www.reddit.com/submit?url=https%3A%2F%2Fwww.engadget.com%2F2019-08-24-crime-allegation-in-space.html%3Fsoc_src%3Dsocial-sh%26soc_trk%3Dreddit&title=Divorce dispute leads to accusation of crime in space" rel="noopener noreferrer" target=_blank title="Share on Reddit" data-ylk=t4:ctrl;sec:consolidated-share-popup;elm:share;elmt:sh-rd;slk:Reddit;itc:0;outcm:share;rspns:op><span class="reddit icon"><svg width=20 height=20 viewBox="0 0 24 24"><path d="M15.018,12.556 C14.252,12.556 13.631,13.194 13.631,13.98 C13.631,14.766 14.252,15.404 15.018,15.404 C15.783,15.404 16.404,14.766 16.404,13.98 C16.404,13.194 15.783,12.556 15.018,12.556 M11.9505,18.2157 C10.4555,18.2157 9.0725,17.5367 8.6595,16.5977 L7.7875,16.9817 C8.3645,18.2887 10.0365,19.1677 11.9795,19.1677 C13.8925,19.1677 15.5655,18.2887 16.1415,16.9817 L15.2705,16.5977 C14.8575,17.5367 13.4735,18.2157 11.9505,18.2157 M8.9425,15.3663 C9.7085,15.3663 10.3295,14.7293 10.3295,13.9423 C10.3295,13.1563 9.7085,12.5193 8.9425,12.5193 C8.1765,12.5193 7.5565,13.1563 7.5565,13.9423 C7.5565,14.7293 8.1765,15.3663 8.9425,15.3663 M20.96,13.641 L20.939,13.652 C20.537,12.343 19.677,11.172 18.485,10.253 C18.876,9.92 19.362,9.729 19.876,9.729 C21.074,9.729 22.049,10.723 22.049,11.943 C22.049,12.678 21.694,13.222 20.96,13.641 M12.002,21.172 C7.412,21.172 3.818,18.591 3.818,15.297 C3.818,11.92 7.489,9.172 12.002,9.172 C16.513,9.172 20.183,11.92 20.183,15.297 C20.183,18.591 16.589,21.172 12.002,21.172 M1.953,11.944 C1.953,10.723 2.928,9.729 4.126,9.729 C4.637,9.729 5.12,9.919 5.515,10.254 C4.324,11.172 3.465,12.343 3.063,13.652 C2.307,13.219 1.953,12.676 1.953,11.944 M19.264,2.843 C20.23,2.843 21.014,3.627 21.014,4.593 C21.014,5.56 20.23,6.344 19.264,6.344 C18.297,6.344 17.514,5.56 17.514,4.593 C17.514,3.627 18.297,2.843 19.264,2.843 M19.876,8.778 C19.069,8.778 18.303,9.088 17.726,9.645 L17.668,9.699 C16.223,8.82 14.433,8.273 12.483,8.194 L13.775,3.841 L16.672,4.587 L16.671,4.591 C16.671,6.025 17.831,7.186 19.263,7.186 C20.695,7.186 21.856,6.025 21.856,4.593 C21.856,3.161 20.695,2 19.263,2 C18.165,2 17.232,2.686 16.854,3.649 L13.122,2.69 L11.491,8.195 C9.551,8.278 7.771,8.825 6.334,9.699 L6.283,9.653 C5.699,9.088 4.933,8.778 4.126,8.778 C2.403,8.778 1,10.198 1,11.944 C1.001,13.025 1.538,13.879 2.624,14.496 L2.85,14.607 C2.822,14.835 2.805,15.066 2.805,15.298 C2.805,19.151 6.844,22.168 12.002,22.168 C17.158,22.168 21.197,19.151 21.197,15.298 C21.197,15.066 21.18,14.835 21.151,14.607 L21.404,14.482 C22.463,13.879 22.999,13.025 23,11.944 C23,10.198 21.599,8.778 19.876,8.778" /></svg></span></a><a class="link caas-button noborder linkedin" href=https://www.linkedin.com/sharing/share-offsite/?url=https%3A%2F%2Fwww.engadget.com%2F2019-08-24-crime-allegation-in-space.html%3Fsoc_src%3Dsocial-sh%26soc_trk%3Dlinkedin rel="noopener noreferrer" target=_blank title="Share on LinkedIn" data-ylk=t4:ctrl;sec:consolidated-share-popup;elm:share;elmt:sh-li;slk:LinkedIn;itc:0;outcm:share;rspns:op><span class="linkedin icon"><svg width=20 height=20 viewBox="0 0 24 24"><path d="M19.208,19.009 L16.315,19.009 L16.315,14.326 C16.315,13.086 16.091,12.345 15.058,12.345 C14.315,12.345 13.757,12.75 13.387,13.138 L13.387,19.009 L10.494,19.009 L10.494,10.399 L13.387,10.399 L13.387,11.549 C13.891,11.005 14.758,10.33 15.971,10.33 C18.002,10.33 19.208,11.639 19.208,14.326 L19.208,19.009 Z M6.965,9.16 C6.043,9.16 5.295,8.412 5.295,7.491 C5.295,6.568 6.043,5.82 6.965,5.82 C7.889,5.82 8.636,6.568 8.636,7.491 C8.636,8.412 7.889,9.16 6.965,9.16 L6.965,9.16 Z M8.43,19.009 L5.537,19.009 L5.537,10.399 L8.43,10.399 L8.43,19.009 Z M22.893,1.727 C22.82,1.565 22.732,1.431 22.624,1.323 C22.517,1.216 22.391,1.135 22.249,1.081 C22.105,1.027 21.944,1 21.766,1 L2.234,1 C2.054,1 1.893,1.027 1.751,1.081 C1.608,1.135 1.483,1.216 1.376,1.323 C1.268,1.431 1.178,1.565 1.107,1.727 C1.035,1.888 1,2.041 1,2.184 L1,21.817 C1,21.961 1.035,22.113 1.107,22.274 C1.178,22.436 1.268,22.57 1.376,22.678 C1.483,22.785 1.608,22.866 1.751,22.92 C1.893,22.974 2.054,23 2.234,23 L21.766,23 C21.944,23 22.105,22.974 22.249,22.92 C22.391,22.866 22.517,22.785 22.624,22.678 C22.732,22.57 22.82,22.436 22.893,22.274 C22.964,22.113 23,21.961 23,21.817 L23,2.184 C23,2.041 22.964,1.888 22.893,1.727 L22.893,1.727 Z" /></svg></span></a></div></div></div><a class="link caas-button caas-tooltip flickrComment caas-comment  top comment-cba6dabc-4fcf-38b7-9338-71b89880a1a8" href=# data-uuid=cba6dabc-4fcf-38b7-9338-71b89880a1a8 title=Reactions data-wf-lightbox-target=#wafer-lightbox-spotIm-template data-wf-lightbox-key=spotim-comments-lightbox data-ylk=t4:ctrl;elm:cmmt_open;elmt:cmnt;slk:comments;itc:1><span class="newComment icon"><svg width=16 height=16 viewBox="0 0 24 24"><path d="M0,4.4C0,2,2,0,4.5,0h15.1C22,0,24,2,24,4.4v10.9c0,2.4-2,4.4-4.5,4.4H6.4l-4.5,4.1C1.6,23.9,1.4,24,1.1,24c-0.1,0-0.3,0-0.4-0.1c-0.4-0.2-0.7-0.6-0.7-1v-7.6V4.4z M19.5,17.5c1.3,0,2.3-1,2.3-2.2V4.4c0-1.2-1-2.2-2.3-2.2H4.5c-1.3,0-2.3,1-2.3,2.2l0,16.2l3.3-3.1H19.5z M17.5,6.5H6.6C6,6.5,5.5,7,5.5,7.6c0,0.6,0.5,1.1,1.1,1.1h10.9c0.6,0,1.1-0.5,1.1-1.1C18.6,7,18.1,6.5,17.5,6.5z M5.5,12c0-0.6,0.5-1.1,1.1-1.1h10.9c0.6,0,1.1,0.5,1.1,1.1c0,0.6-0.5,1.1-1.1,1.1H6.6C6,13.1,5.5,12.6,5.5,12z" style=fill:#9a58b5 /></svg></span><span data-id=cba6dabc-4fcf-38b7-9338-71b89880a1a8 data-type=commentsCount data-site=engadget data-source=spotIm class="reactions-count  caas-dynamic-count ">0</span></a></div></div></div></div><div><div class="caas-carousel caas-cover" role=group aria-label=Slideshow tabindex=0><div class=caas-img-container><div class=caas-carousel-slides><div class=caas-carousel-slide><figure class=caas-carousel-figure role=img aria-label="SLIDE 1 of 2, NASA via AP" style=padding-bottom:67%><img alt="NASA via AP" src=https://s.yimg.com/ny/api/res/1.2/sg6.tLcbW_jQCaBwMJwbnw--/YXBwaWQ9aGlnaGxhbmRlcjt3PTk2MDtoPTY0MA--/https://o.aolcdn.com/images/dims?crop=4500%2C3000%2C0%2C0&quality=85&format=jpg&resize=1600%2C1067&image_uri=https%3A%2F%2Fs.yimg.com%2Fos%2Fcreatr-images%2F2019-08%2F983f1170-c6ae-11e9-9fda-3e6cde7fd289&client=a1acac3e1b3290917d92&signature=37009f7e109fe9cebc9b3b5fb19f678170b53546 class=caas-img data-headline data-credit data-caption="NASA via AP"></figure></div><div class=caas-carousel-slide><figure class=caas-carousel-figure role=img aria-label="SLIDE 2 of 2, NASA via AP" style=padding-bottom:67%><img alt="NASA via AP" data-src=https://s.yimg.com/ny/api/res/1.2/rrr6tgYppOfAypblo_uTnQ--/YXBwaWQ9aGlnaGxhbmRlcjt3PTk2MDtoPTY0MA--/https://o.aolcdn.com/images/dims?crop=4500%2C3000%2C0%2C0&quality=85&format=jpg&resize=1600%2C1067&image_uri=https://s.yimg.com/os/creatr-images/2019-08/983f1170-c6ae-11e9-9fda-3e6cde7fd289&client=a1acac3e1b3290917d92&signature=37009f7e109fe9cebc9b3b5fb19f678170b53546 class=caas-img data-headline="Headline test" data-credit data-caption="NASA via AP"><noscript><img alt="NASA via AP" src=https://s.yimg.com/ny/api/res/1.2/rrr6tgYppOfAypblo_uTnQ--/YXBwaWQ9aGlnaGxhbmRlcjt3PTk2MDtoPTY0MA--/https://o.aolcdn.com/images/dims?crop=4500%2C3000%2C0%2C0&quality=85&format=jpg&resize=1600%2C1067&image_uri=https://s.yimg.com/os/creatr-images/2019-08/983f1170-c6ae-11e9-9fda-3e6cde7fd289&client=a1acac3e1b3290917d92&signature=37009f7e109fe9cebc9b3b5fb19f678170b53546 class=caas-img></noscript></figure></div></div><button class="link prev-button caas-button hide" title="Previous photo" data-ylk=sec:sshow;elm:arrow;outcm:ct;slk:previous;ct:slideshow aria-hidden=true tabindex=-1><span class="prev icon"><svg width=15 height=15 viewBox="0 0 32 32"><path d="M22.72.665c-.886-.887-2.323-.887-3.21 0L4.175 16 19.51 31.335c.442.443 1.023.665 1.604.665s1.162-.222 1.605-.665c.886-.887.886-2.324 0-3.21L10.596 16 22.72 3.876c.887-.887.887-2.324 0-3.21z" style=fill:#ffffff /></svg></span></button><button class="link next-button caas-button hide" title="Next photo" data-ylk=sec:sshow;elm:arrow;outcm:ct;slk:next;ct:slideshow aria-hidden=true tabindex=-1><span class="next icon"><svg width=15 height=15 viewBox="0 0 32 32"><path d="M7.06.665c.887-.887 2.324-.887 3.21 0L25.606 16 10.27 31.335c-.442.443-1.023.665-1.604.665s-1.162-.222-1.605-.665c-.886-.887-.886-2.324 0-3.21L19.185 16 7.06 3.876c-.887-.887-.887-2.324 0-3.21z" style=fill:#ffffff /></svg></span></button></div><div class=figure-meta aria-hidden=true><span class=current>1</span> / 2</div><div class="description collapse-caption" aria-hidden=true><h2 class=description-headline></h2><div class=description-caption-wrapper><span class=description-caption><p>Paragraph test 1</span></div></div></div></div><div class=caas-body><p>p test 2<a href="https://www.engadget.com/2019-04-17-nasa-christina-koch-longest-spaceflight-record.html" data-ylk="slk:spaceflight milestone;elm:context_link;itc:0" class="link ">p test 3</a>p test 4<a href="https://www.nytimes.com/2019/08/23/us/in space.nasa-astronaut-anne-mcclain.html?action=click&amp;module=News&amp;pgtype=Homepage" rel="nofollow noopener" target="_blank" data-ylk="slk:filed complaints;elm:context_link;itc:0" class="link ">filed complaints</a> accusing Worden&#39;s estranged spouse, astronaut Anne McClain (above), of committing a crime while  When McClain appeared to know of Worden&#39;s spending habits despite an ongoing separation battle, Worden found that McClain had accessed their still-linked bank account while aboard the International Space Station -- supposedly committing the crimes of identity theft and improper access to private financial records.</p><p>McClain has acknowledged accessing the account, but said she was checking finances as she&#39;d always done during the relationship, with Worden&#39;s knowledge and permission. McClain was &quot;totally cooperating,&quot; according to her attorney.</p><p>This isn&#39;t the first instance of crime allegations linked to space programs. The <em>New York Times</em> noted that NASA caught a widow trying to sell a Moon rock in 2011, while an Austrian businessman sued a space tourism outfit in 2017 to reclaim his deposit on a trip that appeared to have stalled. However, this appears to be the first instance of accusing an individual of committing a crime while in space. While the subject of crime in space was likely to come up at some point, few would have expected it to come relatively quickly.</p><p>It raises a number of questions about handling accusations of crime in space. Will NASA have to give lawyers access to a secretive network as part of the discovery process, for instance? And while the jurisdiction is relatively clear here (it went through a NASA network and affected someone in the US), it&#39;s not certain what would happen if there was a crime committed exclusively in space. These questions will have to be answered at some point, though, especially when the US intends a <a href="https://www.engadget.com/2019-03-26-vp-pence-vows-us-return-to-moon-by-2024.html" data-ylk="slk:more permanent human presence;elm:context_link;itc:0" class="link ">more permanent human presence</a> in space.</p></div></div></div></div><aside class="caas-aside-section caas-slotted-aside"><div class=caas-non-sticky-sda><div class=caas-slotted-rr-sda><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_top_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"top_right"}}}}' id=sda-LREC-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC2-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC2-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right"}}}}' id=sda-LREC2-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC2-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC3-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC3-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right_2"}}}}' id=sda-LREC3-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC3-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC4-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC4-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right_3"}}}}' id=sda-LREC4-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC4-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC5-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC5-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right_4"}}}}' id=sda-LREC5-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC5-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC6-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC6-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right_5"}}}}' id=sda-LREC6-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC6-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC7-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC7-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right_6"}}}}' id=sda-LREC7-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC7-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC8-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC8-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right_7"}}}}' id=sda-LREC8-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC8-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div><div class=caas-da><div data-wf-benji-page-context='{"url":"https://www.engadget.com/2019-08-24-crime-allegation-in-space.html","spaceid":"1197802876","site":"engadget","hashtag":"spaceflight;nasa;summer-worden;divorce;anne-mcclain;crime;internet;space;gear;legalmatters;science;astronomy","lmsid":"a0V0W00000HOQu8UAH","pt":"content","pd":"non_modal","pct":"story"}' data-wf-benji-wafer-config='{"updateI13n":true}' data-wf-benji-config='{"positions":{"LREC9-cba6dabc-4fcf-38b7-9338-71b89880a1a8":{"id":"LREC9-cba6dabc-4fcf-38b7-9338-71b89880a1a8","inView":"onLoad","path":"/22888152279/us/eng/ros/dt/us_eng_ros_dt_mid_right","region":"index","size":[[300,250],[300,600]],"kvs":{"loc":"mid_right_8"}}}}' id=sda-LREC9-cba6dabc-4fcf-38b7-9338-71b89880a1a8 class="wafer-benji caas-sda-benji-ad" data-wf-trigger=onLoad data-wf-margin="400 0" data-hide-ad-string><div id=LREC9-cba6dabc-4fcf-38b7-9338-71b89880a1a8></div></div></div></div></div></aside></div></div></div></article></div></div><script type=text/javascript>`))
	}))
	defer ts.Close()

	mockClient := new(MockRandomProxyClient)
	scraper := NewScraper(mockClient)

	// Mock the NewHTTPClientWithRandomProxy to return a client that uses the mock server
	mockClient.On("NewHTTPClientWithRandomProxy").Return(&http.Client{}, nil)

	t.Run("successful scraping", func(t *testing.T) {
		// Call the Scrap function with the mock server URL
		result, err := scraper.Scrap(ts.URL, 1)

		assert.NoError(t, err)
		assert.Contains(t, result, "Headline test")
		assert.Contains(t, result, "Paragraph test 1")
		assert.Contains(t, result, "p test 2")
		assert.Contains(t, result, "p test 3")
		assert.Contains(t, result, "p test 4")
		// Additional assertions as needed
	})

	// Add more test cases here for different scenarios like error handling, retry logic, etc.
}

// Additional helper functions for setting up mocks and expected responses can be written here.