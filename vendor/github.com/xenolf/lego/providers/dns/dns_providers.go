// Factory for DNS providers
package dns

import (
	"fmt"

	"github.com/xenolf/lego/acmev2"
	"github.com/xenolf/lego/providers/dns/auroradns"
	"github.com/xenolf/lego/providers/dns/azure"
	"github.com/xenolf/lego/providers/dns/bluecat"
	"github.com/xenolf/lego/providers/dns/cloudflare"
	"github.com/xenolf/lego/providers/dns/cloudxns"
	"github.com/xenolf/lego/providers/dns/digitalocean"
	"github.com/xenolf/lego/providers/dns/dnsimple"
	"github.com/xenolf/lego/providers/dns/dnsmadeeasy"
	"github.com/xenolf/lego/providers/dns/dnspod"
	"github.com/xenolf/lego/providers/dns/duckdns"
	"github.com/xenolf/lego/providers/dns/dyn"
	"github.com/xenolf/lego/providers/dns/exec"
	"github.com/xenolf/lego/providers/dns/exoscale"
	"github.com/xenolf/lego/providers/dns/fastdns"
	"github.com/xenolf/lego/providers/dns/gandi"
	"github.com/xenolf/lego/providers/dns/gandiv5"
	"github.com/xenolf/lego/providers/dns/glesys"
	"github.com/xenolf/lego/providers/dns/godaddy"
	"github.com/xenolf/lego/providers/dns/googlecloud"
	"github.com/xenolf/lego/providers/dns/lightsail"
	"github.com/xenolf/lego/providers/dns/linode"
	"github.com/xenolf/lego/providers/dns/namecheap"
	"github.com/xenolf/lego/providers/dns/namedotcom"
	"github.com/xenolf/lego/providers/dns/ns1"
	"github.com/xenolf/lego/providers/dns/otc"
	"github.com/xenolf/lego/providers/dns/ovh"
	"github.com/xenolf/lego/providers/dns/pdns"
	"github.com/xenolf/lego/providers/dns/rackspace"
	"github.com/xenolf/lego/providers/dns/rfc2136"
	"github.com/xenolf/lego/providers/dns/route53"
	"github.com/xenolf/lego/providers/dns/vultr"
)

func NewDNSChallengeProviderByName(name string) (acmev2.ChallengeProvider, error) {
	var err error
	var provider acmev2.ChallengeProvider
	switch name {
	case "azure":
		provider, err = azure.NewDNSProvider()
	case "auroradns":
		provider, err = auroradns.NewDNSProvider()
	case "bluecat":
		provider, err = bluecat.NewDNSProvider()
	case "cloudflare":
		provider, err = cloudflare.NewDNSProvider()
	case "cloudxns":
		provider, err = cloudxns.NewDNSProvider()
	case "digitalocean":
		provider, err = digitalocean.NewDNSProvider()
	case "dnsimple":
		provider, err = dnsimple.NewDNSProvider()
	case "dnsmadeeasy":
		provider, err = dnsmadeeasy.NewDNSProvider()
	case "dnspod":
		provider, err = dnspod.NewDNSProvider()
	case "duckdns":
		provider, err = duckdns.NewDNSProvider()
	case "dyn":
		provider, err = dyn.NewDNSProvider()
	case "fastdns":
		provider, err = fastdns.NewDNSProvider()
	case "exoscale":
		provider, err = exoscale.NewDNSProvider()
	case "gandi":
		provider, err = gandi.NewDNSProvider()
	case "gandiv5":
		provider, err = gandiv5.NewDNSProvider()
	case "glesys":
		provider, err = glesys.NewDNSProvider()
	case "gcloud":
		provider, err = googlecloud.NewDNSProvider()
	case "godaddy":
		provider, err = godaddy.NewDNSProvider()
	case "lightsail":
		provider, err = lightsail.NewDNSProvider()
	case "linode":
		provider, err = linode.NewDNSProvider()
	case "manual":
		provider, err = acmev2.NewDNSProviderManual()
	case "namecheap":
		provider, err = namecheap.NewDNSProvider()
	case "namedotcom":
		provider, err = namedotcom.NewDNSProvider()
	case "rackspace":
		provider, err = rackspace.NewDNSProvider()
	case "route53":
		provider, err = route53.NewDNSProvider()
	case "rfc2136":
		provider, err = rfc2136.NewDNSProvider()
	case "vultr":
		provider, err = vultr.NewDNSProvider()
	case "ovh":
		provider, err = ovh.NewDNSProvider()
	case "pdns":
		provider, err = pdns.NewDNSProvider()
	case "ns1":
		provider, err = ns1.NewDNSProvider()
	case "otc":
		provider, err = otc.NewDNSProvider()
	case "exec":
		provider, err = exec.NewDNSProvider()
	default:
		err = fmt.Errorf("Unrecognised DNS provider: %s", name)
	}
	return provider, err
}
