package nostr

import (
	"testing"
)

var s = []string{"wss://relay.damus.io", "wss://nostr.mom", "wss://nostr.slothy.win", "wss://relay.stoner.com", "wss://nostr.einundzwanzig.space", "wss://nos.lol", "wss://relay.nostr.band", "wss://relay.oldcity-bitcoiners.info", "wss://nostr.massmux.com", "wss://nostr-relay.schnitzel.world", "wss://relay.nostr.com.au", "wss://knostr.neutrine.com", "wss://nostr.nodeofsven.com", "wss://nostr.vulpem.com", "wss://relay.farscapian.com", "wss://relay.sovereign-stack.org", "wss://relay.lexingtonbitcoin.org", "wss://relay-pub.deschooling.us", "wss://nostr.easydns.ca", "wss://relay.dwadziesciajeden.pl", "wss://nostr.600.wtf", "wss://e.nos.lol", "wss://ragnar-relay.com", "wss://nostr.data.haus", "wss://nostr.wine", "wss://nostr.koning-degraaf.nl", "wss://nostr.cheeserobot.org", "wss://nostr.thank.eu", "wss://relay.hamnet.io", "wss://nostr.blockpower.capital", "wss://nostr.sidnlabs.nl", "wss://nostr.inosta.cc", "wss://nostr21.com", "wss://arc1.arcadelabs.co", "wss://nostr.ch3n2k.com", "wss://relay.nostrview.com", "wss://relay.nostromo.social", "wss://offchain.pub", "wss://relay.nostr.wirednet.jp", "wss://nostr.l00p.org", "wss://lightningrelay.com", "wss://bitcoinmaximalists.online", "wss://private.red.gb.net", "wss://relay.nostrid.com", "wss://nostr.uthark.com", "wss://relay.nostrcheck.me", "wss://nostrelay.yeghro.site", "wss://relay.nostr.vet", "wss://nostr.yuv.al", "wss://nostrue.com", "wss://nostr.danvergara.com", "wss://nproxy.kristapsk.lv", "wss://nostr.topeth.info", "wss://nostr.bitcoiner.social", "wss://relay.orange-crush.com", "wss://nostr.spaceshell.xyz", "wss://nostr.screaminglife.io", "wss://nostr.roundrockbitcoiners.com", "wss://relay.f7z.io", "wss://relay.nostrology.org", "wss://nostr.bch.ninja", "wss://relay.nostrati.com", "wss://nostr-relay.nokotaro.com", "wss://relay.snort.social", "wss://nostr.lu.ke", "wss://atlas.nostr.land", "wss://nostr.fmt.wiz.biz", "wss://global-relay.cesc.trade", "wss://nostr.pjv.me", "wss://relay.roli.social", "wss://brb.io", "wss://eden.nostr.land", "wss://nostr-verified.wellorder.net", "wss://nostr.noones.com", "wss://relay.nostr.nu", "wss://nostr-relay.bitcoin.ninja", "wss://deschooling.us", "wss://freespeech.casa", "wss://bitcoiner.social", "wss://nostr.1f52b.xyz", "wss://nostr.sebastix.dev", "wss://relay-verified.deschooling.us", "wss://nostr.cizmar.net", "wss://n.wingu.se", "wss://relay.nostr.hach.re", "wss://nostr.bitcoinplebs.de", "wss://nostr.corebreach.com", "wss://tmp-relay.cesc.trade", "wss://nostr.mutinywallet.com", "wss://nostr.decentony.com", "wss://nostr.tools.global.id", "wss://xmr.usenostr.org", "wss://nostr.naut.social", "wss://at.nostrworks.com", "wss://nostr.sovbit.host", "wss://nostr.1sat.org", "wss://mastodon.cloud/api/v1/streaming", "wss://nostr.256k1.dev", "wss://relay.beta.fogtype.com", "wss://rsslay.ch3n2k.com", "wss://nostr.rocketnode.space", "wss://relay.nostr.bg", "wss://nostr.malin.onl", "wss://nostr.globals.fans", "wss://nostr.primz.org", "wss://relay.johnnyasantos.com", "wss://btc.klendazu.com", "wss://relay.n057r.club", "wss://slick.mjex.me", "wss://nostr.lorentz.is", "wss://relay.nostrified.org", "wss://relay.primal.net", "wss://nostr.cercatrova.me", "wss://nostr.swiss-enigma.ch", "wss://relay.honk.pw", "wss://nostr-relay.derekross.me", "wss://puravida.nostr.land", "wss://nostr.sectiontwo.org", "wss://nostr.oxtr.dev", "wss://relay.s3x.social", "wss://eosla.com", "wss://nostr.liberty.fans", "wss://nostr.cro.social", "wss://nostrrelay.com", "wss://nostr-pub.semisol.dev", "wss://nostr.semisol.dev", "wss://misskey.io", "wss://nostr.pobblelabs.org", "wss://relay.nostr.wf", "wss://nostr.land", "wss://btcpay.kukks.org/nostr/ws", "wss://relay.mostr.pub", "wss://relay.nostrplebs.com", "wss://purplepag.es", "wss://yestr.me", "wss://relayable.org", "wss://paid.nostrified.org", "wss://nostr-02.dorafactory.org", "wss://nostr.zbd.gg", "wss://relay.hodl.ar", "wss://relay.nostr.sc", "wss://feeds.nostr.band/nostrhispano", "wss://nostr.portemonero.com", "wss://search.nos.today", "wss://relay.minds.com/nostr/v1/ws", "wss://welcome.nostr.wine", "wss://yabu.me", "wss://adult.18plus.social", "wss://nrelay.c-stellar.net", "wss://nostrja-kari.heguro.com", "wss://therelayofallrelays.nostr1.com", "wss://nostr-relay.app", "wss://rly.nostrkid.com", "wss://nostr.filmweb.pl", "wss://nostr.openhoofd.nl", "wss://nostr.strits.dk", "wss://relay.nostr.moctane.com", "wss://nostr.thesamecat.io", "wss://relay.deezy.io", "wss://nostr-usa.ka1gbeoa21bnm.us-west-2.cs.amazonlightsail.com", "wss://relay.poster.place", "wss://freerelay.xyz", "wss://nostr.mining.sc", "wss://nostr.xmr.rocks", "wss://nostr.dvdt.dev", "wss://paid.nostr.lc", "wss://ithurtswhenip.ee", "wss://nerostr.xmr.rocks", "wss://powrelay.xyz", "wss://relay.rebelbase.site", "wss://relay.vanderwarker.family", "wss://wc1.current.ninja", "wss://relay.toastr.space", "wss://nostr.ingwie.me", "wss://nostrpub.yeghro.site", "wss://relay2.nostrchat.io", "wss://relay1.nostrchat.io", "wss://relay.devstr.org", "wss://pleb.cloud", "wss://nostr.dakukitsune.ca", "wss://nostr.hifish.org", "wss://nostr2.sanhauf.com", "wss://nostrja-kari-nip50.heguro.com", "wss://nostrua.com", "wss://nostr.otc.sh", "wss://relay.nsecbunker.com", "wss://relay.nostr.youlot.org", "wss://nostr.hekster.org", "wss://relay.layer.systems", "wss://nostr.schorsch.fans", "wss://nostr.reelnetwork.eu", "wss://nostr.hexhex.online", "wss://relay.kisiel.net.pl", "wss://relay.nostr.directory", "wss://relay.wavlake.com", "wss://nostr-tbd.website", "wss://nostr.438b.net", "wss://relay.nostr.lighting", "wss://nostr.sagaciousd.com", "wss://nostr.donky.social", "wss://nostr.metamadeenah.com", "wss://kukks.org/nostr", "wss://nostrja-world-relays-test.heguro.com", "wss://multiplextr.coracle.social", "wss://relay.mutinywallet.com", "wss://nostril.cam", "wss://nostr.btc-library.com", "wss://relay.getalby.com/v1", "wss://rss.nos.social", "wss://nostr.overmind.lol", "wss://beta.nostril.cam", "wss://dev.nostrplayground.com", "wss://nostr-01.yakihonne.com", "wss://nostr.fort-btc.club", "wss://relay.bitcoinpark.com", "wss://nostr.olwe.link", "wss://nostr01.counterclockwise.io", "wss://relap.orzv.workers.dev", "wss://christpill.nostr1.com", "wss://relay.verified-nostr.com", "wss://nostr.sathoarder.com", "wss://nostr.sudocarlos.com", "wss://wbc.nostr1.com", "wss://relay.protest.net", "wss://lnbits.eldamar.icu/nostrrelay/relay", "wss://pater.nostr1.com", "wss://nostr.heliodex.cf", "wss://nostr.getgle.org", "wss://nostr.lnbitcoin.cz", "wss://lnbits.michaelantonfischer.com/nostrrelay/michaelantonf", "wss://nostr.relayable.org", "wss://relay.casualcrypto.date", "wss://lightning.benchodroff.com/nostrclient/api/v1/relay", "wss://relay.notmandatory.org", "wss://nostr.cx.ms", "wss://fiatjaf.com", "wss://relay.despera.space", "wss://bitstack.app", "wss://nostr-relay.psfoundation.info", "wss://relay.hrf.org", "wss://relay.orangepill.ovh", "wss://relay.relayable.org", "wss://nostr.rubberdoll.cc", "wss://relay.nostrassets.com", "wss://relay.ingwie.me", "wss://soloco.nl", "wss://nostr.dlsouza.lol", "wss://relay.kamp.site", "wss://nostr.heavyrubberslave.com", "wss://relay.keychat.io", "wss://he.relayable.org", "wss://au.relayable.org", "wss://relay.ohbe.me", "wss://nostr.lecturify.net", "wss://nostr.bitcoinist.org", "wss://nostr.stakey.net", "wss://a.nos.lol", "wss://eu.purplerelay.com", "wss://greensoul.space", "wss://21ideas.nostr1.com", "wss://hotrightnow.nostr1.com", "wss://relay.bitcoinbarcelona.xyz", "wss://rly.bopln.com", "wss://nostr.fractalized.net", "wss://r.hostr.cc", "wss://relay.stemstr.app", "wss://bevo.nostr1.com", "wss://cache2.primal.net/v1", "wss://relay.nquiz.io", "wss://nostr.self-determined.de", "wss://gardn.nostr1.com", "wss://feedstr.nostr1.com", "wss://nostr.zoz-serv.org", "wss://nostr.girino.org", "wss://nostr.gerbils.online", "wss://nostr.novacisko.cz", "wss://nostr.0x7e.xyz", "wss://sakhalin.nostr1.com", "wss://adre.su", "wss://remnant.cloud", "wss://nostr.kungfu-g.rip", "wss://relay.guggero.org", "wss://relay.minibolt.info", "wss://cache1.primal.net/v1", "wss://misskey.art", "wss://merrcurrup.railway.app", "wss://nostr-dev.zbd.gg", "wss://relay.oke.minds.io/nostr/v1/ws", "wss://tw.purplerelay.com", "wss://za.purplerelay.com", "wss://novoa.nagoya", "wss://submarin.online", "wss://misskey.04.si", "wss://social.camph.net", "wss://in.purplerelay.com", "wss://nostr.hashi.sbs", "wss://nostr.atitlan.io", "wss://nostr.dbtc.link", "wss://relay.0v0.social", "wss://nsrelay.assilvestrar.club", "wss://nostr.lbdev.fun", "wss://relay.corpum.com", "wss://relay.benthecarman.com", "wss://relay.nostrhub.fr", "wss://relay.noswhere.com", "wss://orangepiller.org", "wss://relay.satoshidnc.com", "wss://relay.siamstr.com", "wss://relay.camelus.app", "wss://frens.nostr1.com", "wss://datagrave.wild-vibes.ts.net/nostr", "wss://ae.purplerelay.com", "wss://kr.purplerelay.com", "wss://vitor.nostr1.com", "wss://chefstr.nostr1.com", "wss://n.ok0.org", "wss://nostr.8777.ch", "wss://prism.nostr1.com", "wss://sfr0.nostr1.com", "wss://us.nostr.land", "wss://relay.nostr.jabber.ch", "wss://testnet.plebnet.dev/nostrrelay/1", "wss://th2.nostr.earnkrub.xyz", "wss://relay.nostr.pt", "wss://studio314.nostr1.com", "wss://relay.piazza.today", "wss://relay.mattybs.lol", "wss://relay.exit.pub", "wss://nostr.plantroon.com", "wss://hk.purplerelay.com", "wss://nostr.se7enz.com", "wss://au.purplerelay.com", "wss://br.purplerelay.com", "wss://ca.purplerelay.com", "wss://ch.purplerelay.com", "wss://nostr.jcloud.es", "wss://cl.purplerelay.com", "wss://us.nostr.wine", "wss://relay.proxymana.net", "wss://cellar.nostr.wine", "wss://fl.purplerelay.com", "wss://nostr.stupleb.cc", "wss://frjosh.nostr1.com", "wss://inbox.nostr.wine", "wss://relay.minibits.cash", "wss://ftp.halifax.rwth-aachen.de/nostr", "wss://nostr-02.yakihonne.com", "wss://nostpy.serverless-nostr.com", "wss://shawn.nostr1.com", "wss://relay.gems.xyz", "wss://nostr.hubmaker.io", "wss://loli.church", "wss://lnbits.aruku.kro.kr/nostrrelay/private", "wss://sushi.ski", "wss://xmr.ithurtswhenip.ee", "wss://relay.nos.social", "wss://obiurgator.thewhall.com", "wss://relay.staging.geyser.fund", "wss://nostr.neilalexander.dev", "wss://relay.causes.com", "wss://nostr.psychoet.nexus", "wss://relay.artx.market", "wss://nostr.tavux.tech", "wss://nostr-03.dorafactory.org", "wss://lnbits.btconsulting.nl/nostrrelay/nostr", "wss://nostr.zoel.network", "wss://relay.geyser.fund", "wss://lnbits.papersats.io/nostrclient/api/v1/relay", "wss://creatr.nostr.wine", "wss://riray.nostr1.com", "wss://pyramid.fiatjaf.com", "wss://node.coincreek.com/nostrclient/api/v1/relay", "wss://relay.tunestr.io", "wss://ren.nostr1.com", "wss://cache0.primal.net/cache17", "wss://nostr.33co.de", "wss://nostr.t-rg.ws", "wss://theforest.nostr1.com", "wss://de.purplerelay.com", "wss://nostrdevs.nostr1.com", "wss://njump.me", "wss://nostr.cahlen.org", "wss://strfry.chatbett.de", "wss://milwaukie.nostr1.com", "wss://relay.bitmapstr.io", "wss://directory.yabu.me", "wss://noxir.fly.dev", "wss://srtrelay.c-stellar.net", "wss://nostr.reckless.dev", "wss://super-relay.cesc.trade", "wss://bucket.coracle.social", "wss://nostr.lopp.social", "wss://relay.notoshi.win", "wss://sfgbsfg431512asf124as.xyz", "wss://relay.hawties.xyz", "wss://jp.purplerelay.com", "wss://lnbits.satoshibox.io/nostrclient/api/v1/relay", "wss://rkgk.moe", "wss://relay.dev.bdw.to", "wss://ksdfgsg12412312sdgfq23.online", "wss://gjmhmhgi789hjgdyerysergdfvbncte.fyi", "wss://carlos-cdb.top", "wss://relay.moinsen.com", "wss://hayloo88.nostr1.com", "wss://a214g24132sa2fas354411f234125.xyz", "wss://itchy-goldenrod-furconthophagus.scarab.im", "wss://140.f7z.io", "wss://noxir.kpherox.dev", "wss://taipei.scarab.im", "wss://nostr.madco.me", "wss://jingle.carlos-cdb.top", "wss://staging.yabu.me", "wss://nostrrelay.win", "wss://adfasfasfadsdfasfasf3123412ewfas.xyz", "wss://nostr.sats.li", "wss://nostr.notribe.net", "wss://paid.relay.vanderwarker.family", "wss://nostr.animeomake.com", "wss://23asdfasf2r341gnbrrhjhwggadffgasfsadfasafa.xyz", "wss://287avuahggadsd213rg18aga3yg3whg8g8afg.xyz", "wss://rnostr.onrender.com", "wss://u42qujakzujtggyjtdcjl6b582z1gvckh52.xyz", "wss://relay.famstr.xyz", "wss://relay.roygbiv.guide", "wss://nostr.ussenterprise.xyz", "wss://relay3.blogstr.app", "wss://nostr.searx.is", "wss://xxmmrr.shogatsu.ovh", "wss://nostr.extrabits.io", "wss://relay.magiccity.live", "wss://dev-relay.nostrassets.com", "wss://nostr.happytavern.co", "wss://relay.agorist.space", "wss://nostr.intrepid18.com", "wss://nostr.ser1.net", "wss://marmot.nostr1.com", "wss://island.nostr1.com", "wss://nostr.jfischer.org", "wss://relay.earthly.land", "wss://bostr.nokotaro.work", "wss://relay.angor.io", "wss://r314y.0xd43m0n.xyz", "wss://privateisland.club", "wss://relay.nostrosity.com", "wss://bostr.nokotaro.com", "wss://relay.lawallet.ar", "wss://relay.nimo.cash", "wss://relay.nostr.net", "wss://relay.cosmicbolt.net", "wss://nostr.dmgd.monster", "wss://nostr-relay.jezza.online", "wss://bouncer.nostree.me", "wss://relay.lax1dude.net", "wss://relay.13room.space", "wss://relay.westernbtc.com", "wss://nostr.char0n.ch", "wss://misskey.cloud", "wss://relay.blackbyte.nl", "wss://reactions.v0l.io", "wss://relay.lumina.rocks", "wss://relay2.denostr.com", "wss://fiatjaf.nostr1.com", "wss://nostr.at", "wss://nostr.babyshark.win", "wss://nostr.coincrowd.fund", "wss://relay.nostar.org", "wss://nostr.jaonoctus.dev", "wss://bouncer.minibolt.info", "wss://relay.s-w.art", "wss://nostr.petrkr.net/strfry", "wss://bits.lnbc.sk/nostrclient/api/v1/relay", "wss://thenewregi.me/nostr", "wss://echo.websocket.org", "wss://dmux.net/apps/strfry", "wss://relay.nip05.social", "wss://dash.mockingyou.com", "wss://relay.gasteazi.net", "wss://mm.suzuqi.com", "wss://relay.highlighter.com", "wss://relay.urbanzap.space", "wss://relay.nostrasia.net", "wss://strfry.iris.to", "wss://dreamofthe90s.nostr1.com", "wss://relay.dolu.dev", "wss://relay2.nostrasia.net", "wss://dev-relay.kube.b-n.space", "wss://bostr.bitcointxoko.com", "wss://kmc-nostr.amiunderwater.com", "wss://relay.usefusion.ai", "wss://nostr.bit4use.com", "wss://nostr.coinfund.app", "wss://nostr.tbai.me:592", "wss://nostr.lifeonbtc.xyz", "wss://relay.nostpy.lol", "wss://hist.nostr.land", "wss://primal-cache.mutinywallet.com/v1", "wss://freelay.sovbit.host", "wss://relay.nsec.app", "wss://cl4.tnix.dev", "wss://public-relay.swallow.solutions", "wss://nostr.yuhr.org", "wss://nostr.myshosholoza.co.za", "wss://nostr.btczh.tw", "wss://nostr.sebastix.social", "wss://galaxy13.nostr1.com", "wss://chorus.pjv.me", "wss://nostr.palandi.cloud", "wss://nostr.millonair.us", "wss://nostr.codingarena.top", "wss://nostr.pistaum.com", "wss://relay.wikifreedia.xyz", "wss://strfry.openhoofd.nl", "wss://nostria.space", "wss://nostr.a2x.pub", "wss://notes.miguelalmodo.com", "wss://support.nostr1.com", "wss://relay.phantom-power.coracle.tools", "wss://relay.varke.eu", "wss://relay.cxplay.org", "wss://freedomhub1.reelnetwork.eu", "wss://blowater.nostr1.com", "wss://relay.refinery.coracle.tools", "wss://nostr.drafted.pro", "wss://nostr.lnwallet.app", "wss://relays.diggoo.com", "wss://relay.nostrainsley.coracle.tools", "wss://nostr.hashbang.nl", "wss://relay.zone667.com", "wss://bostr.freespeech.casa", "wss://custom.fiatjaf.com", "wss://nostr-relay.algotech.io", "wss://bostr.online", "wss://relay.momostr.pink", "wss://nosflare.plebes.fans", "wss://multiplexer.huszonegy.world", "wss://misskey.design", "wss://hodlbod.coracle.tools", "wss://ir.purplerelay.com", "wss://us.purplerelay.com", "wss://me.purplerelay.com", "wss://countries.fiatjaf.com", "wss://zap.nostr1.com", "wss://avb.nostr1.com", "wss://relay.monstr.ing", "wss://tigs.nostr1.com", "wss://nostr.lnpixels.com", "wss://nostr-relay.philipcristiano.com", "wss://svjgceikshv.sylonet.io", "wss://nosflare.royalgarter.workers.dev", "wss://relay.martien.io", "wss://nostr-relay.cbrx.io", "wss://relay.neuance.net", "wss://invillage-outvillage.com", "wss://clnbits.diynodes.com/nostrclient/api/v1/relay", "wss://relay.oh-happy-day.xyz", "wss://nostr.skitso.business", "wss://history.nostr.watch", "wss://relaypag.es", "wss://nostr.greenfield.xyz", "wss://nr.rosano.ca", "wss://nostr.n-ti.me", "wss://relay2.angor.io", "wss://bostr.lightningspore.com", "wss://relay.satsdays.com", "wss://relay.nostreggs.io", "wss://longhorn.bgp.rodeo", "wss://nostr.stubby.dev", "wss://junxingwang.org", "wss://nostrrelay.stewlab.win", "wss://nostr.manasiwibi.com", "wss://nostr.goldfigure.ai", "wss://core.btcmap.org/nostrrelay/relay", "wss://nostr.flamingo-mail.com", "wss://relay-jp.haniwar.com", "wss://nostr.bitcoinvn.io", "wss://relay.tagayasu.xyz", "wss://cfrelay.haorendashu.workers.dev", "wss://lunchbox.sandwich.farm", "wss://relay.bitdevs.tw", "wss://nostr.falci.me", "wss://cryptowolf.nostr1.com", "wss://khatru-test.puhcho.me", "wss://nostr.javi.space", "wss://nostr.cloud.vinney.xyz", "wss://relay.arrakis.lat", "wss://relay.illuminodes.com", "wss://relay.test.nquiz.io", "wss://relay.bostr.online", "wss://relay.satlantis.io", "wss://relay.unknown.cloud", "wss://relay.twixy.app", "wss://nostr.brackrat.com", "wss://relay1.xfire.to:", "wss://relay.vantis.ninja", "wss://nostr.kolbers.de", "wss://relay.thoreau.fyi", "wss://nostr.cybercan.click", "wss://rnostr.fractalized.net", "wss://nostr2.fractalized.net", "wss://need-venues-rules-identifying.trycloudflare.com", "wss://auth.nostr1.com", "wss://relay.hex3f.com", "wss://nostr.uneu.net", "wss://nostr.daedaluslabs.io", "wss://nostr1.daedaluslabs.io", "wss://nostr2.daedaluslabs.io", "wss://sydney.nostr21.net", "wss://nostr3.daedaluslabs.io", "wss://kiwibuilders.nostr21.net", "wss://cfrelay.puhcho.workers.dev", "wss://nosdrive.app/relay", "wss://relay.nostr.cymru", "wss://ditto.fractalized.net/relay", "wss://nostr-relay.aaroniumii.com", "wss://relay.livefreebtc.com", "wss://relay.fountain.fm", "wss://moonboi.nostr1.com", "wss://nostr.pailakapo.com", "wss://nostr.sovrgn.co.za", "wss://relay.alex71btc.com", "wss://nostr.kloudcover.com", "wss://relay.bohe.lol", "wss://pl.czas.xyz", "wss://cfrelay.polok.workers.dev", "wss://deno.blowater.app", "wss://nostr.dodge.me.uk", "wss://relay.braydon.com", "wss://nostr.mdip.yourself.dev", "wss://us.nostr.fractalized.net", "wss://nostrich.adagio.tw", "wss://nostr.workehub.com", "wss://rp.czas.xyz", "wss://relay.gnostr.cloud", "wss://nostr.whipq.com", "wss://relay.magnifiq.tech", "wss://r.mleku.net", "wss://nostr.15b.blue", "wss://af.purplerelay.com", "wss://sos.l.test.denostr.com", "wss://sos.m.test.denostr.com", "wss://sos.s.test.denostr.com", "wss://nostr.netbros.com", "wss://relay.staging.flashapp.me", "wss://strfry-test-external.summary.news", "wss://hodla.nu/nostrrelay/1", "wss://libretechsystems.nostr1.com", "wss://relay.rootservers.eu", "wss://nostr.polonkai.eu", "wss://dvms.f7z.io", "wss://nostr.thurk.org", "wss://strfry.corebreach.com", "wss://afraid-simulations-collecting-mastercard.trycloudflare.com", "wss://ivy.hnslogin.world", "wss://unhostedwallet.com", "wss://thecitadel.nostr1.com", "wss://red.nostr-relay-bouncer.xyz", "wss://guaporelay.aaroniumii.com"}

func TestNamedLock(t *testing.T) {
	for _, url := range s {
		unlock := namedLock(url)
		unlock()
	}
}