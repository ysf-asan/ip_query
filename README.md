
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

_Adım 1:_ Go'yu Yükleyin

Eğer Go dilini bilgisayarınızda kurulu değilse, Go'nun resmi web sitesinden yükleyebilirsiniz.

_Adım 2:_ Projeyi İndirin

Proje dosyasını bilgisayarınıza indirin:

```sh
git clone https://https://github.com/ysf-asan/ip_query
cd ip_query
```

_Adım 3:_ Bağımlılıkları Yükleyin

Bu projede HTTP istekleri ve JSON işleme için Go'nun dahili kütüphaneleri kullanılmıştır, bu nedenle ek bir bağımlılık yüklemenize gerek yoktur.

_Adım 4:_ Uygulamayı Çalıştırın

Proje dizininde şu komutu çalıştırarak uygulamayı başlatabilirsiniz:
```sh
go run main.go
```
Uygulama çalıştıktan sonra, IP adresinizle ilgili bilgileri terminal üzerinden görebilirsiniz.

__Kullanım__

Çalıştırdığınızda uygulama, aşağıdaki bilgileri döndürecektir:

IP Adresi: Kullanıcının cihazına ait IP adresi.

Şehir: IP adresinin bağlı olduğu şehir.

Ülke: IP adresinin bağlı olduğu ülke.

Organizasyon: IP adresinin ait olduğu organizasyon (İSS veya şirket).

Konum: IP adresinin coğrafi konum bilgileri (enlem ve boylam).


__API Kullanımı__

Bu proje, ipinfo.io API'sini kullanarak IP adresi bilgilerini alır. Ücretsiz bir API anahtarı ile günde sınırlı sayıda sorgu yapılabilir. Daha fazla sorgu yapmak için bir API anahtarı alabilir ve projede kullanabilirsiniz.

__Katkı__

Projeye katkıda bulunmak için pull request gönderebilirsiniz. Herhangi bir öneri veya hata bildirimi için Issues kısmını kullanabilirsiniz.

__Lisans__

Bu proje MIT Lisansı ile lisanslanmıştır.
