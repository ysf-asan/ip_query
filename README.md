
IP Adresi Sorgulama Aracı

Bu Go dilinde yazılmış küçük bir uygulama, kullanıcının cihazının IP adresini sorgular ve IP ile ilgili çeşitli bilgileri döndürür. Bu bilgiler arasında IP adresi, şehir, ülke, organizasyon ve konum yer alır.

__Özellikler__

Kullanıcının cihazının IP adresini alır.

IP adresine ait şehir, ülke, organizasyon ve konum bilgilerini getirir.

Basit ve hızlı kullanım için HTTP tabanlı ipinfo.io API'sini kullanır.


__Gereksinimler__

Go 1.18 veya daha yeni bir sürümü.

İnternet bağlantısı (API'ye istek gönderebilmek için).


__Kurulum__

Adım 1: Go'yu Yükleyin

Eğer Go dilini bilgisayarınızda kurulu değilse, Go'nun resmi web sitesinden yükleyebilirsiniz.

Adım 2: Projeyi İndirin

Proje dosyasını bilgisayarınıza indirin:

```sh
git clone https://github.com/yourusername/ip-address-lookup.git
cd ip-address-lookup
```

Adım 3: Bağımlılıkları Yükleyin

Bu projede HTTP istekleri ve JSON işleme için Go'nun dahili kütüphaneleri kullanılmıştır, bu nedenle ek bir bağımlılık yüklemenize gerek yoktur.

Adım 4: Uygulamayı Çalıştırın

Proje dizininde şu komutu çalıştırarak uygulamayı başlatabilirsiniz:
```sh
go run main.go
```
Uygulama çalıştıktan sonra, IP adresinizle ilgili bilgileri terminal üzerinden görebilirsiniz.

Kullanım

Çalıştırdığınızda uygulama, aşağıdaki bilgileri döndürecektir:

IP Adresi: Kullanıcının cihazına ait IP adresi.

Şehir: IP adresinin bağlı olduğu şehir.

Ülke: IP adresinin bağlı olduğu ülke.

Organizasyon: IP adresinin ait olduğu organizasyon (İSS veya şirket).

Konum: IP adresinin coğrafi konum bilgileri (enlem ve boylam).


API Kullanımı

Bu proje, ipinfo.io API'sini kullanarak IP adresi bilgilerini alır. Ücretsiz bir API anahtarı ile günde sınırlı sayıda sorgu yapılabilir. Daha fazla sorgu yapmak için bir API anahtarı alabilir ve projede kullanabilirsiniz.

Katkı

Projeye katkıda bulunmak için pull request gönderebilirsiniz. Herhangi bir öneri veya hata bildirimi için Issues kısmını kullanabilirsiniz.

Lisans

Bu proje MIT Lisansı ile lisanslanmıştır.
