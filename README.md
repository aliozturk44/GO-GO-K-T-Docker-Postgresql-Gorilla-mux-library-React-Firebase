Bu projede golang dili ile bir çok kütüphane ve framework kullanarak bir mikroservis mimarisi inşa ettim. 
Ön yüz olarak React kullandım.
Veritabanı olarak postgresql ve firebase kullandım. 
Firebase kullanma sebebim cloud veri tabanı ile yüklenen resimlere ulaşabilmek.
Ancak firebase veritabanının süresi bittiğinde resimlere ulaşamayabilirsiniz o kısım çalışmayabilir.
Bu sebeple yeni bir firebase veritabanı kurup bağlarsanız o kısım da sizde sorunsuz çalışacaktır.
Ve son olarak da docker ile tüm projeyi ayağa kaldırdım. tabiki siz isterseniz go run main.go ve npm start komutlarını kullanarak back endi ve ön yüzü ayrı ayrı da ayağa kaldırabilirsiniz.
