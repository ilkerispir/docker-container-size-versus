FROM ruby:alpine

WORKDIR /usr/src/app
COPY Gemfile Gemfile.lock ./
ENV BUNDLE_FROZEN=true
RUN gem install bundler && bundle install --without test

COPY . ./

CMD ["ruby", "./app.rb"]