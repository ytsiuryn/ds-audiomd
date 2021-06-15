# ds-audiomd

Модуль описывает структуры хранения метаданных аудио, их заполнение, контроль и
оптимизацию для хранения и экспорта в JSON.

Используется для информационного обмена метаданными между микросервисами
[MdReader](https://github.com/ytsiuryn/ds-mdreader), [Discogs](https://github.com/ytsiuryn/ds-discogs) и [Musicbrainz](https://github.com/ytsiuryn/ds-musicbrainz) и проч.

Для *хранения метаданных* верхнеуровневой точкой представления является **релиз** (Release).

Для *представления предложений метаданных от клиентов внешних информационных сервисов*
([Discogs](https://www.discogs.com/developers), [Musicbrainz](https://musicbrainz.org/doc/MusicBrainz_API)) **предложение** (Suggestion).

Модуль также содержит необходимую функциональность для оценки соответствия предложений
online-сервисов имеющимся метаданным.