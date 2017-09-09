package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Rendering Delimited Blocks", func() {
	It("source block with multiple lines", func() {
		content := "```\nsome source code\n\nhere\n```"
		expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
</div>
</div>`
		verify(GinkgoT(), expected, content)
	})
})