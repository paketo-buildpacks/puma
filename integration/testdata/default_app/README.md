# Puma default web application

Run with:

```
$ pack build pumaapp -b gcr.io/paketo-buildpacks/mri -b gcr.io/paketo-buildpacks/bundler \
-b gcr.io/paketo-buildpacks/bundle-install -b gcr.io/paketo-buildpacks/puma

$ docker run --rm -it -p 9292:9292 pumaapp
```

Enjoy:

```
$ curl 0.0.0.0:9292
```

### Notes

By default, if no configuration file is specified, Puma will look for a
configuration file at `config/puma.rb`.
See https://www.rubydoc.info/gems/puma/3.6.2#configuration-file
