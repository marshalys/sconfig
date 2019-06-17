package sconfig_test

import (
	"testing"

	"github.com/marshalys/sconfig"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SConfigTestSuite struct {
	suite.Suite
	config sconfig.SConfig
}

func (s *SConfigTestSuite) SetupSuite() {
	s.config = sconfig.New()
	const configPath = "./testdata/test.yml"
	err := s.config.LoadConfig(configPath)
	t := s.T()
	assert.Equal(t, nil, err)
}

func (s *SConfigTestSuite) TestGet() {
	t := s.T()
	val, ok := s.config.Get("context")
	_ = assert.Equal(t, true, ok) && assert.NotEqual(t, nil, val)
	val, ok = s.config.Get("context.timeoutMilliseconds")
	_ = assert.Equal(t, true, ok) && assert.NotEqual(t, nil, val)
	val, ok = s.config.Get("notexist")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, nil, val)
}

func (s *SConfigTestSuite) TestGetString() {
	t := s.T()
	val, ok := s.config.GetString("services.goodsService.address")
	_ = assert.Equal(t, true, ok) && assert.Equal(t, "127.0.0.1:50001", val)
	val, ok = s.config.GetString("context.timeoutMilliseconds")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, "", val)
	val, ok = s.config.GetString("notexist")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, "", val)
}

func (s *SConfigTestSuite) TestGetBool() {
	t := s.T()
	val, ok := s.config.GetBool("upgrade.auto")
	_ = assert.Equal(t, true, ok) && assert.Equal(t, true, val)
	val, ok = s.config.GetBool("upgrade.weekly")
	_ = assert.Equal(t, true, ok) && assert.Equal(t, false, val)
	val, ok = s.config.GetBool("context.timeoutMilliseconds")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, false, val)
	val, ok = s.config.GetBool("notexist")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, false, val)
}

func (s *SConfigTestSuite) TestGetInt() {
	t := s.T()
	val, ok := s.config.GetInt("context.timeoutMilliseconds")
	_ = assert.Equal(t, true, ok) && assert.Equal(t, 3000, val)
	val, ok = s.config.GetInt("upgrade.times")
	_ = assert.Equal(t, true, ok) && assert.Equal(t, -1, val)
	val, ok = s.config.GetInt("upgrade.auto")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, 0, val)
	val, ok = s.config.GetInt("notexist")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, 0, val)
}

func (s *SConfigTestSuite) TestGetFloat64() {
	t := s.T()
	val, ok := s.config.GetFloat64("upgrade.float64")
	_ = assert.Equal(t, true, ok) && assert.Equal(t, 1111111111111111111111.0, val)
	val, ok = s.config.GetFloat64("context.timeoutMilliseconds")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, 0.0, val)
	val, ok = s.config.GetFloat64("upgrade.auto")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, 0.0, val)
	val, ok = s.config.GetFloat64("notexist")
	_ = assert.Equal(t, false, ok) && assert.Equal(t, 0.0, val)
}

func (s *SConfigTestSuite) TestGetStringSlice() {
	t := s.T()
	val, ok := s.config.GetStringSlice("safeHosts")
	_ = assert.Equal(t, true, ok) && assert.Equal(t, 3, len(val)) && assert.Equal(t, "admin-dev.test.com", val[0])
	val, ok = s.config.GetStringSlice("context.timeoutMilliseconds")
	_ = assert.Equal(t, false, ok) && assert.True(t, val == nil)
	val, ok = s.config.GetStringSlice("upgrade.auto")
	_ = assert.Equal(t, false, ok) && assert.True(t, val == nil)
	val, ok = s.config.GetStringSlice("notexist")
	_ = assert.Equal(t, false, ok) && assert.True(t, val == nil)
}

func (s *SConfigTestSuite) TestAllSettings() {
	t := s.T()
	val := s.config.AllSettings()
	_ = assert.True(t, val != nil)
}

type cacheSetting struct {
	DefaultExpirationSeconds uint32
	CleanupIntervalSeconds   uint32
}

type corsSetting struct {
	AllowOrigins []string
}

func (s *SConfigTestSuite) TestUnmarshalKey() {
	t := s.T()
	cache := &cacheSetting{}
	err := s.config.UnmarshalKey("cache.memory", cache)
	_ = assert.Equal(t, nil, err) && assert.Equal(t, uint32(3600), cache.DefaultExpirationSeconds)
	cors := &corsSetting{}
	err = s.config.UnmarshalKey("cors", cors)
	_ = assert.Equal(t, nil, err) && assert.Equal(t, "http://admin-dev.test.com", cors.AllowOrigins[0])
}

func TestSConfigTestSuite(t *testing.T) {
	suite.Run(t, new(SConfigTestSuite))
}
